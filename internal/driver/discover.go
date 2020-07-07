//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"encoding/binary"
	"fmt"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	edgexModels "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.impcloud.net/RSP-Inventory-Suite/device-llrp-go/internal/llrp"
	"io"
	"math/bits"
	"net"
	"regexp"
	"sync"
	"time"
)

// todo: most of these need to be configurable
const (
	probeTimeout          = 1 * time.Second
	scanVirtualInterfaces = false
	profileName           = "Device.LLRP.Profile"
	probeAsyncLimit       = 1000
)

var (
	llrpPortStr = "5084" // todo: support TLS connections
)

// virtualRegex is a regular expression to determine if an interface is likely to be a virtual interface
var virtualRegex = regexp.MustCompile("^(?:docker[0-9]+|br-.*|virbr[0-9]+.*|docker_gwbridge|veth.*)$")

// autoDiscover probes all addresses in the local network to attempt to discover any possible
// RFID readers that support LLRP.
func autoDiscover() (discovered []dsModels.DiscoveredDevice) {
	ipCh := make(chan uint32, 5*probeAsyncLimit)
	done := make(chan struct{})
	resultCh := make(chan uint32)
	var wg sync.WaitGroup

	deviceMap := makeDeviceMap()

	// start this before adding any ips so they are ready to process
	for i := 0; i < probeAsyncLimit; i++ {
		go ipWorker(deviceMap, &wg, done, ipCh, resultCh)
	}

	discovered = make([]dsModels.DiscoveredDevice, 0)
	go resultsWorker(resultCh, &discovered)

	netCh := getIPv4Nets(scanVirtualInterfaces)
	ipGenerator(&wg, netCh, ipCh)

	wg.Wait()
	close(ipCh)
	close(resultCh)
	close(done)

	return discovered
}

// makeDeviceMap creates a lookup table of existing devices in order to skip scanning
func makeDeviceMap() map[string]bool {
	devices := driver.service().Devices()
	deviceMap := make(map[string]bool, len(devices))

	for _, d := range devices {
		tcpInfo := d.Protocols["tcp"]
		if tcpInfo == nil {
			log.Infof("found registered device without tcp protocol information: %s", d.Name)
			continue
		}
		host := tcpInfo["host"]
		port := tcpInfo["port"]
		if host == "" || port == "" {
			log.Warnf("registered device is missing required tcp protocol information %s: %v", d.Name, d.Protocols)
			continue
		}

		deviceMap[makeDeviceName(host, port)] = true
	}

	return deviceMap
}

func getIPv4Nets(includeVirtual bool) <-chan *net.IPNet {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Error(err)
		return nil
	}

	out := make(chan *net.IPNet, len(ifaces))

	go func() {
		for _, iface := range ifaces {
			if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
				continue
			}

			addrs, err := iface.Addrs()
			if err != nil {
				log.Error(err)
				continue
			}

			if !includeVirtual && virtualRegex.MatchString(iface.Name) {
				driver.lc.Debug("Skipping virtual network interface: " + iface.Name)
				continue
			}

			for _, addr := range addrs {
				if inet, ok := addr.(*net.IPNet); ok {
					if inet.IP.To4() == nil {
						continue
					}
					driver.lc.Info(fmt.Sprintf("Scan interface %s: %v", iface.Name, inet))
					out <- inet
				}
			}
		}
		close(out)
	}()
	return out
}

func ipGenerator(wg *sync.WaitGroup, netCh <-chan *net.IPNet, ipCh chan<- uint32) {
	var addr net.IP
	for inet := range netCh {
		addr = inet.IP.To4()
		if addr == nil {
			continue
		}

		mask := inet.Mask
		if len(mask) == net.IPv6len {
			mask = mask[12:]
		} else if len(mask) != net.IPv4len {
			continue
		}

		umask := binary.BigEndian.Uint32(mask)
		maskSz := bits.OnesCount32(umask)
		if maskSz <= 1 {
			continue // skip point-to-point connections
		} else if maskSz >= 31 {
			// special cases where only 1 ip is valid in subnet
			wg.Add(1)
			ipCh <- binary.BigEndian.Uint32(inet.IP)
			continue
		}

		netId := binary.BigEndian.Uint32(addr) & umask // network ID
		bcast := netId ^ (^umask)
		for ip := netId + 1; ip < bcast; ip++ {
			if netId&umask != ip&umask {
				continue
			}

			wg.Add(1)
			ipCh <- ip
		}
	}
}

// probe attempts to make a connection to a specific ip and port to determine
// if an LLRP reader exists at that network address
func probe(ip string, port string) error {
	addr := ip + ":" + port
	//driver.lc.Debug(fmt.Sprintf("probe: %s", ip))
	conn, err := net.DialTimeout("tcp", addr, probeTimeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	driver.lc.Info("connection dialed")

	buf := make([]byte, llrp.HeaderSz)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		log.Error(err)
		return err
	}

	h := llrp.Header{}
	err = h.UnmarshalBinary(buf)
	log.Debugf("connection header: %+v", h)

	if h.Type() != llrp.ReaderEventNotification {
		return errors.Errorf("expected %v, got %v", llrp.ReaderEventNotification, h.Type())
	}

	driver.lc.Info("Reader successfully discovered @ " + addr)
	return nil
}

func resultsWorker(results <-chan uint32, discovered *[]dsModels.DiscoveredDevice) {
	ip := net.IP([]byte{0, 0, 0, 0})
	for a := range results {
		binary.BigEndian.PutUint32(ip, a)
		*discovered = append(*discovered, newDiscoveredDevice(ip.String(), llrpPortStr))
	}
}

// ipWorker pulls uint32s, convert to IPs, and sends back successful probes
func ipWorker(deviceMap map[string]bool, wg *sync.WaitGroup, done <-chan struct{}, ipCh <-chan uint32, results chan<- uint32) {
	ip := net.IP([]byte{0, 0, 0, 0})
	for {
		select {
		case a, ok := <-ipCh:
			if !ok {
				return
			}
			binary.BigEndian.PutUint32(ip, a)

			ipStr := ip.String()
			if _, found := deviceMap[makeDeviceName(ipStr, llrpPortStr)]; found {
				log.Infof("Skip scan of %s, device already registered", ipStr)
				wg.Done()
				continue
			}

			if err := probe(ipStr, llrpPortStr); err == nil {
				results <- a
			}
			wg.Done()
		case <-done:
			return
		}
	}
}

func makeDeviceName(ip string, port string) string {
	return ip + "_" + port
}

// newDiscoveredDevice takes the host and port number of a discovered LLRP reader and prepares it for
// registering use with EdgeX
func newDiscoveredDevice(ip string, port string) dsModels.DiscoveredDevice {
	deviceName := makeDeviceName(ip, port)

	return dsModels.DiscoveredDevice{
		Name: deviceName,
		Protocols: map[string]edgexModels.ProtocolProperties{
			"tcp": {
				"host": ip,
				"port": port,
				"llrp": "llrp",
			},
		},
		Description: "LLRP RFID Reader",
		Labels:      nil,
	}
}