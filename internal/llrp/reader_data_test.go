//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package llrp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestReader_withRecordedData(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		msg := f.Name()[:len(f.Name())-len(".json")]
		t.Run(msg, compareMessages(msg))
	}
}

func compareMessages(msgName string) func(t *testing.T) {
	type binRoundTrip interface {
		UnmarshalBinary(data []byte) error
		MarshalBinary() ([]byte, error)
	}
	var v binRoundTrip
	switch "Msg" + msgName {
	case MsgGetReaderCapabilitiesResponse.String():
		v = &GetReaderCapabilitiesResponse{}
	case MsgGetReaderConfigResponse.String():
		v = &GetReaderConfigResponse{}
	case MsgGetAccessSpecsResponse.String():
		v = &GetAccessSpecsResponse{}
	case MsgGetROSpecsResponse.String():
		v = &GetROSpecsResponse{}
	case MsgCloseConnectionResponse.String():
		v = &CloseConnectionResponse{}
	}

	// This tests the following two conversions using data captured from a reader:
	//   JSON -> Go -> binary & check it matches original binary
	// binary -> Go -> JSON   & check it matches original JSON
	return func(t *testing.T) {
		if v == nil {
			t.Fatalf("unknown message type: %s", msgName)
		}

		var jsonData, byteData, marshaledBin, marshaledJSON []byte
		var err error

		// load data files
		jsonData, err = ioutil.ReadFile("testdata/" + msgName + ".json")
		if err != nil {
			t.Fatalf("can't read .json file: %v", err)
		}

		byteData, err = ioutil.ReadFile("testdata/" + msgName + ".bytes")
		if err != nil {
			t.Fatalf("can't read .bytes file: %v", err)
		}

		// unmarshal original JSON form
		if err = json.Unmarshal(jsonData, v); err != nil {
			t.Fatal(err)
		}

		// marshal resulting struct to binary
		if marshaledBin, err = v.MarshalBinary(); err != nil {
			t.Fatal(err)
		}

		// confirm binary matches original
		checkBytesEq(t, byteData, marshaledBin)

		// get a new v (so we're not duplicating list items)
		v = reflect.New(reflect.TypeOf(v).Elem()).Interface().(binRoundTrip)

		// unmarshal original binary form to struct
		if err = v.UnmarshalBinary(marshaledBin); err != nil {
			t.Fatal(err)
		}

		// marshal struct back to JSON
		marshaledJSON, err = json.MarshalIndent(v, "", "\t")
		if err != nil {
			t.Fatal(err)
		}

		// confirm JSON data matches original
		if !checkJSONEq(t, jsonData, marshaledJSON) {
			t.Logf("%s", marshaledJSON)
		}
	}
}

// checkJSONEq checks that two json data byte arrays are equal.
//
// If not, it prints a side-by-side diff around the first difference,
// replacing tabs with 4 periods.
//
// Returns true if the two arrays are equal; false otherwise.
func checkJSONEq(t *testing.T, jsonData, marshaled []byte) (matched bool) {
	t.Helper()

	lines1, lines2 := bytes.Split(jsonData, []byte("\n")), bytes.Split(marshaled, []byte("\n"))

	matched = len(lines1) == len(lines2)
	smaller := len(lines2)
	if len(lines1) < len(lines2) {
		smaller = len(lines1)
		matched = false
	}

	firstDiff := 0
	for ; firstDiff < smaller; firstDiff++ {
		if !bytes.Equal(lines1[firstDiff], lines2[firstDiff]) {
			matched = false
			break
		}
	}

	if matched {
		return
	}

	const contextAbove = 4
	start := firstDiff - contextAbove
	if start < 0 {
		start = 0
	}

	const contextTotal = 16
	end := start + contextTotal
	if end > len(lines1)-1 {
		end = len(lines1) - 1
	}

	if end > len(lines2)-1 {
		end = len(lines2) - 1
	}

	longest := 1
	for i := start; i < end; i++ {
		lines1[i] = bytes.ReplaceAll(lines1[i], []byte("\t"), []byte("...."))
		lines2[i] = bytes.ReplaceAll(lines2[i], []byte("\t"), []byte("...."))
		if len(lines1[i]) > longest {
			longest = len(lines1[i])
		}
		if len(lines2[i]) > longest {
			longest = len(lines2[i])
		}
	}

	diff := bytes.Buffer{}
	fmt.Fprintf(&diff, "%-[1]*s  |  %s\n", longest, "    --Original JSON Data--", "    --Marshaled Result--")
	for i := start; i < end; i++ {
		if i == firstDiff {
			msg := "--first diff below this line--"
			fmt.Fprintf(&diff, "%[1]*s\n", longest+3+len(msg)/2, msg)
		}
		fmt.Fprintf(&diff, "%-[1]*s  |  %s\n", longest, lines1[i], lines2[i])
	}

	t.Errorf("JSON data mismatched; first difference around line %d:\n%s",
		firstDiff, diff.String())
	return
}

func checkBytesEq(t *testing.T, b1, b2 []byte) (matched bool) {
	t.Helper()

	matched = len(b1) == len(b2)
	smaller := len(b2)
	if len(b1) < len(b2) {
		smaller = len(b1)
		matched = false
	}

	firstDiff := 0
	for ; firstDiff < smaller; firstDiff++ {
		if b1[firstDiff] != b2[firstDiff] {
			matched = false
			break
		}
	}

	if matched {
		return
	}

	start := firstDiff - 4
	if start < 0 {
		start = 0
	}

	end := start + 8
	if end > len(b1)-1 {
		end = len(b1) - 1
	}

	if end > len(b2)-1 {
		end = len(b2) - 1
	}

	t.Errorf("byte data mismatched starting at byte %d; surrounding bytes:\n"+
		"%# 02x\n"+
		"%# 02x", firstDiff, b1[start:end], b2[start:end])
	return
}
