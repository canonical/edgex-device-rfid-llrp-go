//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package llrp

import (
	"reflect"
	"testing"
)

// Test Message 46, GetSupportedVersion.
func TestGetSupportedVersion_roundTrip(t *testing.T) {
	m := GetSupportedVersion{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetSupportedVersion
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 56, GetSupportedVersionResponse.
func TestGetSupportedVersionResponse_roundTrip(t *testing.T) {
	m := GetSupportedVersionResponse{
		CurrentVersion:      Version1_0_1,
		MaxSupportedVersion: Version1_0_1,
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetSupportedVersionResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 47, SetProtocolVersion.
func TestSetProtocolVersion_roundTrip(t *testing.T) {
	m := SetProtocolVersion{
		TargetVersion: Version1_0_1,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 SetProtocolVersion
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 57, SetProtocolVersionResponse.
func TestSetProtocolVersionResponse_roundTrip(t *testing.T) {
	m := SetProtocolVersionResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 SetProtocolVersionResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 1, GetReaderCapabilities.
func TestGetReaderCapabilities_roundTrip(t *testing.T) {
	m := GetReaderCapabilities{
		ReaderCapabilitiesRequestedData: ReaderCapLLRPCapabilities,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetReaderCapabilities
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 11, GetReaderCapabilitiesResponse.
func TestGetReaderCapabilitiesResponse_roundTrip(t *testing.T) {
	m := GetReaderCapabilitiesResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetReaderCapabilitiesResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 20, AddROSpec.
func TestAddROSpec_roundTrip(t *testing.T) {
	m := AddROSpec{
		ROSpec: ROSpec{
			ROSpecID:           2147483648,
			Priority:           128,
			ROSpecCurrentState: ROSpecStateInactive,
			ROBoundarySpec: ROBoundarySpec{
				StartTrigger: ROSpecStartTrigger{
					Trigger: ROStartTriggerPeriodic,
				},
				StopTrigger: ROSpecStopTrigger{
					Trigger:              ROStopTriggerDuration,
					DurationTriggerValue: uint32(2147483648),
				},
			},
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 AddROSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 30, AddROSpecResponse.
func TestAddROSpecResponse_roundTrip(t *testing.T) {
	m := AddROSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 AddROSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 21, DeleteROSpec.
func TestDeleteROSpec_roundTrip(t *testing.T) {
	m := DeleteROSpec{
		ROSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DeleteROSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 31, DeleteROSpecResponse.
func TestDeleteROSpecResponse_roundTrip(t *testing.T) {
	m := DeleteROSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DeleteROSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 22, StartROSpec.
func TestStartROSpec_roundTrip(t *testing.T) {
	m := StartROSpec{
		ROSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 StartROSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 32, StartROSpecResponse.
func TestStartROSpecResponse_roundTrip(t *testing.T) {
	m := StartROSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 StartROSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 23, StopROSpec.
func TestStopROSpec_roundTrip(t *testing.T) {
	m := StopROSpec{
		ROSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 StopROSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 33, StopROSpecResponse.
func TestStopROSpecResponse_roundTrip(t *testing.T) {
	m := StopROSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 StopROSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 24, EnableROSpec.
func TestEnableROSpec_roundTrip(t *testing.T) {
	m := EnableROSpec{
		ROSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 EnableROSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 34, EnableROSpecResponse.
func TestEnableROSpecResponse_roundTrip(t *testing.T) {
	m := EnableROSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 EnableROSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 25, DisableROSpec.
func TestDisableROSpec_roundTrip(t *testing.T) {
	m := DisableROSpec{
		ROSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DisableROSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 35, DisableROSpecResponse.
func TestDisableROSpecResponse_roundTrip(t *testing.T) {
	m := DisableROSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DisableROSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 26, GetROSpecs.
func TestGetROSpecs_roundTrip(t *testing.T) {
	m := GetROSpecs{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetROSpecs
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 36, GetROSpecsResponse.
func TestGetROSpecsResponse_roundTrip(t *testing.T) {
	m := GetROSpecsResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetROSpecsResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 40, AddAccessSpec.
func TestAddAccessSpec_roundTrip(t *testing.T) {
	m := AddAccessSpec{
		AccessSpec: AccessSpec{
			AccessSpecID:  2147483648,
			AntennaID:     32768,
			AirProtocolID: AirProtoEPCGlobalClass1Gen2,
			IsActive:      true,
			ROSpecID:      2147483648,
			Trigger: AccessSpecStopTrigger{
				Trigger:             AccessSpecStopTriggerOperationCount,
				OperationCountValue: 32768,
			},
			AccessCommand: AccessCommand{
				C1G2TagSpec: C1G2TagSpec{
					TagPattern1: C1G2TargetTag{
						C1G2MemoryBank:     uint8(1),
						MatchFlag:          true,
						MostSignificantBit: 32768,
						TagMaskNumBits:     32,
						TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
						TagDataNumBits:     32,
						TagData:            []byte{0x0, 0x1, 0x2, 0x3},
					},
				},
			},
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 AddAccessSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 50, AddAccessSpecResponse.
func TestAddAccessSpecResponse_roundTrip(t *testing.T) {
	m := AddAccessSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 AddAccessSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 41, DeleteAccessSpec.
func TestDeleteAccessSpec_roundTrip(t *testing.T) {
	m := DeleteAccessSpec{
		AccessSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DeleteAccessSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 51, DeleteAccessSpecResponse.
func TestDeleteAccessSpecResponse_roundTrip(t *testing.T) {
	m := DeleteAccessSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DeleteAccessSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 42, EnableAccessSpec.
func TestEnableAccessSpec_roundTrip(t *testing.T) {
	m := EnableAccessSpec{
		AccessSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 EnableAccessSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 52, EnableAccessSpecResponse.
func TestEnableAccessSpecResponse_roundTrip(t *testing.T) {
	m := EnableAccessSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 EnableAccessSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 43, DisableAccessSpec.
func TestDisableAccessSpec_roundTrip(t *testing.T) {
	m := DisableAccessSpec{
		AccessSpecID: 2147483648,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DisableAccessSpec
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 53, DisableAccessSpecResponse.
func TestDisableAccessSpecResponse_roundTrip(t *testing.T) {
	m := DisableAccessSpecResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 DisableAccessSpecResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 44, GetAccessSpecs.
func TestGetAccessSpecs_roundTrip(t *testing.T) {
	m := GetAccessSpecs{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetAccessSpecs
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 54, GetAccessSpecsResponse.
func TestGetAccessSpecsResponse_roundTrip(t *testing.T) {
	m := GetAccessSpecsResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetAccessSpecsResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 45, ClientRequestOp.
func TestClientRequestOp_roundTrip(t *testing.T) {
	m := ClientRequestOp{
		TagReportData: TagReportData{
			EPCData: EPCData{
				EPCNumBits: 32,
				EPC:        []byte{0x0, 0x1, 0x2, 0x3},
			},
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 ClientRequestOp
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 55, ClientRequestOpResponse.
func TestClientRequestOpResponse_roundTrip(t *testing.T) {
	m := ClientRequestOpResponse{
		ClientRequestResponse: ClientRequestResponse{
			AccessSpecID: 2147483648,
			EPCData: EPCData{
				EPCNumBits: 32,
				EPC:        []byte{0x0, 0x1, 0x2, 0x3},
			},
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 ClientRequestOpResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 60, GetReport.
func TestGetReport_roundTrip(t *testing.T) {
	m := GetReport{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetReport
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 61, ROAccessReport.
func TestROAccessReport_roundTrip(t *testing.T) {
	m := ROAccessReport{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 ROAccessReport
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 62, KeepAlive.
func TestKeepAlive_roundTrip(t *testing.T) {
	m := KeepAlive{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 KeepAlive
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 72, KeepAliveAck.
func TestKeepAliveAck_roundTrip(t *testing.T) {
	m := KeepAliveAck{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 KeepAliveAck
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 63, ReaderEventNotification.
func TestReaderEventNotification_roundTrip(t *testing.T) {
	m := ReaderEventNotification{
		ReaderEventNotificationData: ReaderEventNotificationData{
			UTCTimestamp: UTCTimestamp(uint64(9223372036854775808)),
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 ReaderEventNotification
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 64, EnableEventsAndReports.
func TestEnableEventsAndReports_roundTrip(t *testing.T) {
	m := EnableEventsAndReports{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 EnableEventsAndReports
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 100, ErrorMessage.
func TestErrorMessage_roundTrip(t *testing.T) {
	m := ErrorMessage{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 ErrorMessage
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 2, GetReaderConfig.
func TestGetReaderConfig_roundTrip(t *testing.T) {
	m := GetReaderConfig{
		AntennaID:     32768,
		RequestedData: ReaderConfReqAccessReportSpec,
		GPIPortNum:    32768,
		GPOPortNum:    32768,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetReaderConfig
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 12, GetReaderConfigResponse.
func TestGetReaderConfigResponse_roundTrip(t *testing.T) {
	m := GetReaderConfigResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 GetReaderConfigResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 3, SetReaderConfig.
func TestSetReaderConfig_roundTrip(t *testing.T) {
	m := SetReaderConfig{
		ResetToFactoryDefaults: true,
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 SetReaderConfig
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 13, SetReaderConfigResponse.
func TestSetReaderConfigResponse_roundTrip(t *testing.T) {
	m := SetReaderConfigResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 SetReaderConfigResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 14, CloseConnection.
func TestCloseConnection_roundTrip(t *testing.T) {
	m := CloseConnection{}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 CloseConnection
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 4, CloseConnectionResponse.
func TestCloseConnectionResponse_roundTrip(t *testing.T) {
	m := CloseConnectionResponse{
		LLRPStatus: LLRPStatus{
			Status:           StatusMsgMsgUnexpected,
			ErrorDescription: "some arbitrary string",
		},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 CloseConnectionResponse
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Message 1023, CustomMessage.
func TestCustomMessage_roundTrip(t *testing.T) {
	m := CustomMessage{
		VendorID:       2147483648,
		MessageSubtype: 128,
		Data:           []byte{128, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
	}
	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var m2 CustomMessage
	if err := m2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, m, m2)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, m, m2)
	}
}

// Test Parameter 1, AntennaID.
func TestAntennaID_roundTrip(t *testing.T) {
	p := AntennaID(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AntennaID
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 2, FirstSeenUTC.
func TestFirstSeenUTC_roundTrip(t *testing.T) {
	p := FirstSeenUTC(uint64(9223372036854775808))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FirstSeenUTC
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 3, FirstSeenUptime.
func TestFirstSeenUptime_roundTrip(t *testing.T) {
	p := FirstSeenUptime(uint64(9223372036854775808))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FirstSeenUptime
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 4, LastSeenUTC.
func TestLastSeenUTC_roundTrip(t *testing.T) {
	p := LastSeenUTC(uint64(9223372036854775808))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 LastSeenUTC
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 5, LastSeenUptime.
func TestLastSeenUptime_roundTrip(t *testing.T) {
	p := LastSeenUptime(uint64(9223372036854775808))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 LastSeenUptime
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 6, PeakRSSI.
func TestPeakRSSI_roundTrip(t *testing.T) {
	p := PeakRSSI(int8(-1))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 PeakRSSI
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 7, ChannelIndex.
func TestChannelIndex_roundTrip(t *testing.T) {
	p := ChannelIndex(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ChannelIndex
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 8, TagSeenCount.
func TestTagSeenCount_roundTrip(t *testing.T) {
	p := TagSeenCount(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 TagSeenCount
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 9, ROSpecID.
func TestROSpecID_roundTrip(t *testing.T) {
	p := ROSpecID(2147483648)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROSpecID
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 10, InventoryParameterSpecID.
func TestInventoryParameterSpecID_roundTrip(t *testing.T) {
	p := InventoryParameterSpecID(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 InventoryParameterSpecID
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 11, C1G2CRC.
func TestC1G2CRC_roundTrip(t *testing.T) {
	p := C1G2CRC(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2CRC
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 12, C1G2PC.
func TestC1G2PC_roundTrip(t *testing.T) {
	p := C1G2PC{
		EPCMemoryLength: uint8(16),
		HasUserMemory:   true,
		HasXPC:          true,
		IsISO15961:      true,
		AttributesOrAFI: 128,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2PC
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 13, EPC96.
func TestEPC96_roundTrip(t *testing.T) {
	p := EPC96{
		EPC: []byte{128, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 EPC96
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 14, SpecIndex.
func TestSpecIndex_roundTrip(t *testing.T) {
	p := SpecIndex(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 SpecIndex
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 15, ClientRequestOpSpecResult.
func TestClientRequestOpSpecResult_roundTrip(t *testing.T) {
	p := ClientRequestOpSpecResult(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ClientRequestOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 16, AccessSpecID.
func TestAccessSpecID_roundTrip(t *testing.T) {
	p := AccessSpecID(2147483648)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AccessSpecID
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 17, OpSpecID.
func TestOpSpecID_roundTrip(t *testing.T) {
	p := OpSpecID(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 OpSpecID
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 18, C1G2SingulationDetails.
func TestC1G2SingulationDetails_roundTrip(t *testing.T) {
	p := C1G2SingulationDetails{
		NumCollisionSlots: 32768,
		NumEmptySlots:     32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2SingulationDetails
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 19, C1G2XPCW1.
func TestC1G2XPCW1_roundTrip(t *testing.T) {
	p := C1G2XPCW1(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2XPCW1
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 20, C1G2XPCW2.
func TestC1G2XPCW2_roundTrip(t *testing.T) {
	p := C1G2XPCW2(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2XPCW2
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 128, UTCTimestamp.
func TestUTCTimestamp_roundTrip(t *testing.T) {
	p := UTCTimestamp(uint64(9223372036854775808))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 UTCTimestamp
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 129, Uptime.
func TestUptime_roundTrip(t *testing.T) {
	p := Uptime(uint64(9223372036854775808))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 Uptime
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 137, GeneralDeviceCapabilities.
func TestGeneralDeviceCapabilities_roundTrip(t *testing.T) {
	p := GeneralDeviceCapabilities{
		MaxSupportedAntennas:    32768,
		CanSetAntennaProperties: true,
		HasUTCClock:             true,
		DeviceManufacturer:      2147483648,
		Model:                   2147483648,
		FirmwareVersion:         "some arbitrary string",
		ReceiveSensitivities: []ReceiveSensitivityTableEntry{
			{
				Index:              32768,
				ReceiveSensitivity: uint16(64),
			},
		},
		GPIOCapabilities: GPIOCapabilities{
			NumGPIs: 32768,
			NumGPOs: 32768,
		},
		PerAntennaAirProtocols: []PerAntennaAirProtocol{
			{
				AntennaID:      32768,
				AirProtocolIDs: []AirProtocolIDType{AirProtoEPCGlobalClass1Gen2, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2},
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 GeneralDeviceCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 139, ReceiveSensitivityTableEntry.
func TestReceiveSensitivityTableEntry_roundTrip(t *testing.T) {
	p := ReceiveSensitivityTableEntry{
		Index:              32768,
		ReceiveSensitivity: uint16(64),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ReceiveSensitivityTableEntry
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 140, PerAntennaAirProtocol.
func TestPerAntennaAirProtocol_roundTrip(t *testing.T) {
	p := PerAntennaAirProtocol{
		AntennaID:      32768,
		AirProtocolIDs: []AirProtocolIDType{AirProtoEPCGlobalClass1Gen2, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2, AirProtoUnspecified, AirProtoEPCGlobalClass1Gen2},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 PerAntennaAirProtocol
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 141, GPIOCapabilities.
func TestGPIOCapabilities_roundTrip(t *testing.T) {
	p := GPIOCapabilities{
		NumGPIs: 32768,
		NumGPOs: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 GPIOCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 142, LLRPCapabilities.
func TestLLRPCapabilities_roundTrip(t *testing.T) {
	p := LLRPCapabilities{
		CanDoRFSurvey:                          true,
		CanReportBufferFillWarning:             true,
		SupportsClientRequestOpSpec:            true,
		CanDoTagInventoryStateAwareSingulation: true,
		SupportsEventsAndReportHolding:         true,
		MaxPriorityLevelSupported:              128,
		ClientRequestedOpSpecTimeout:           32768,
		MaxROSpecs:                             2147483648,
		MaxSpecsPerROSpec:                      2147483648,
		MaxInventoryParameterSpecsPerAISpec:    2147483648,
		MaxAccessSpecs:                         2147483648,
		MaxOpSpecsPerAccessSpec:                2147483648,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 LLRPCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 143, RegulatoryCapabilities.
func TestRegulatoryCapabilities_roundTrip(t *testing.T) {
	p := RegulatoryCapabilities{
		CountryCode:            Unspecified,
		CommunicationsStandard: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RegulatoryCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 144, UHFBandCapabilities.
func TestUHFBandCapabilities_roundTrip(t *testing.T) {
	p := UHFBandCapabilities{
		TransmitPowerLevels: []TransmitPowerLevelTableEntry{
			{
				Index:              32768,
				TransmitPowerValue: uint16(32768),
			},
		},
		FrequencyInformation: FrequencyInformation{
			Hopping: true,
		},
		C1G2RFModes: UHFC1G2RFModeTable{
			UHFC1G2RFModeTableEntries: []UHFC1G2RFModeTableEntry{
				{
					ModeID:                2147483648,
					DivideRatio:           DRSixtyFourToThree,
					IsEPCHagConformant:    true,
					Modulation:            Miller4,
					ForwardLinkModulation: SingleSidebandASK,
					SpectralMask:          SpectralMaskMultiInterrogator,
					BackscatterDataRate:   uint32(300000),
					PIERatio:              1500,
					MinTariTime:           uint32(9375),
					MaxTariTime:           uint32(9375),
					StepTariTime:          uint32(9375),
				},
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 UHFBandCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 145, TransmitPowerLevelTableEntry.
func TestTransmitPowerLevelTableEntry_roundTrip(t *testing.T) {
	p := TransmitPowerLevelTableEntry{
		Index:              32768,
		TransmitPowerValue: uint16(32768),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 TransmitPowerLevelTableEntry
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 146, FrequencyInformation.
func TestFrequencyInformation_roundTrip(t *testing.T) {
	p := FrequencyInformation{
		Hopping: true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FrequencyInformation
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 147, FrequencyHopTable.
func TestFrequencyHopTable_roundTrip(t *testing.T) {
	p := FrequencyHopTable{
		HopTableID:  128,
		Frequencies: []Kilohertz{uint32(2147483648), uint32(1), uint32(2), uint32(3)},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FrequencyHopTable
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 148, FixedFrequencyTable.
func TestFixedFrequencyTable_roundTrip(t *testing.T) {
	p := FixedFrequencyTable{
		Frequencies: []Kilohertz{uint32(2147483648), uint32(1), uint32(2), uint32(3)},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FixedFrequencyTable
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 149, PerAntennaReceiveSensitivityRange.
func TestPerAntennaReceiveSensitivityRange_roundTrip(t *testing.T) {
	p := PerAntennaReceiveSensitivityRange{
		AntennaID:                  32768,
		ReceiveSensitivityIndexMin: 32768,
		ReceiveSensitivityIndexMax: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 PerAntennaReceiveSensitivityRange
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 177, ROSpec.
func TestROSpec_roundTrip(t *testing.T) {
	p := ROSpec{
		ROSpecID:           2147483648,
		Priority:           128,
		ROSpecCurrentState: ROSpecStateInactive,
		ROBoundarySpec: ROBoundarySpec{
			StartTrigger: ROSpecStartTrigger{
				Trigger: ROStartTriggerPeriodic,
			},
			StopTrigger: ROSpecStopTrigger{
				Trigger:              ROStopTriggerDuration,
				DurationTriggerValue: uint32(2147483648),
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 178, ROBoundarySpec.
func TestROBoundarySpec_roundTrip(t *testing.T) {
	p := ROBoundarySpec{
		StartTrigger: ROSpecStartTrigger{
			Trigger: ROStartTriggerPeriodic,
		},
		StopTrigger: ROSpecStopTrigger{
			Trigger:              ROStopTriggerDuration,
			DurationTriggerValue: uint32(2147483648),
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROBoundarySpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 179, ROSpecStartTrigger.
func TestROSpecStartTrigger_roundTrip(t *testing.T) {
	p := ROSpecStartTrigger{
		Trigger: ROStartTriggerPeriodic,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROSpecStartTrigger
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 180, PeriodicTriggerValue.
func TestPeriodicTriggerValue_roundTrip(t *testing.T) {
	p := PeriodicTriggerValue{
		Offset: uint32(2147483648),
		Period: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 PeriodicTriggerValue
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 181, GPITriggerValue.
func TestGPITriggerValue_roundTrip(t *testing.T) {
	p := GPITriggerValue{
		Port:    32768,
		Event:   true,
		Timeout: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 GPITriggerValue
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 182, ROSpecStopTrigger.
func TestROSpecStopTrigger_roundTrip(t *testing.T) {
	p := ROSpecStopTrigger{
		Trigger:              ROStopTriggerDuration,
		DurationTriggerValue: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROSpecStopTrigger
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 183, AISpec.
func TestAISpec_roundTrip(t *testing.T) {
	p := AISpec{
		AntennaIDs: []AntennaID{32768, 1, 2, 3},
		StopTrigger: AISpecStopTrigger{
			Trigger:              AIStopTriggerGPI,
			DurationTriggerValue: uint32(2147483648),
		},
		InventoryParameterSpecs: []InventoryParameterSpec{
			{
				InventoryParameterSpecID: 32768,
				AirProtocolID:            AirProtoEPCGlobalClass1Gen2,
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AISpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 184, AISpecStopTrigger.
func TestAISpecStopTrigger_roundTrip(t *testing.T) {
	p := AISpecStopTrigger{
		Trigger:              AIStopTriggerGPI,
		DurationTriggerValue: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AISpecStopTrigger
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 185, TagObservationTrigger.
func TestTagObservationTrigger_roundTrip(t *testing.T) {
	p := TagObservationTrigger{
		Trigger:          TagObsTriggerNAttempts,
		NumberofTags:     32768,
		NumberofAttempts: 32768,
		T:                uint16(32768),
		Timeout:          uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 TagObservationTrigger
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 186, InventoryParameterSpec.
func TestInventoryParameterSpec_roundTrip(t *testing.T) {
	p := InventoryParameterSpec{
		InventoryParameterSpecID: 32768,
		AirProtocolID:            AirProtoEPCGlobalClass1Gen2,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 InventoryParameterSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 187, RFSurveySpec.
func TestRFSurveySpec_roundTrip(t *testing.T) {
	p := RFSurveySpec{
		AntennaID:      32768,
		StartFrequency: uint32(2147483648),
		EndFrequency:   uint32(2147483648),
		Trigger: RFSurveySpecStopTrigger{
			Trigger:  RFSurveyStopTriggerDuration,
			Duration: uint32(2147483648),
			N:        uint32(2147483648),
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFSurveySpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 188, RFSurveySpecStopTrigger.
func TestRFSurveySpecStopTrigger_roundTrip(t *testing.T) {
	p := RFSurveySpecStopTrigger{
		Trigger:  RFSurveyStopTriggerDuration,
		Duration: uint32(2147483648),
		N:        uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFSurveySpecStopTrigger
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 207, AccessSpec.
func TestAccessSpec_roundTrip(t *testing.T) {
	p := AccessSpec{
		AccessSpecID:  2147483648,
		AntennaID:     32768,
		AirProtocolID: AirProtoEPCGlobalClass1Gen2,
		IsActive:      true,
		ROSpecID:      2147483648,
		Trigger: AccessSpecStopTrigger{
			Trigger:             AccessSpecStopTriggerOperationCount,
			OperationCountValue: 32768,
		},
		AccessCommand: AccessCommand{
			C1G2TagSpec: C1G2TagSpec{
				TagPattern1: C1G2TargetTag{
					C1G2MemoryBank:     uint8(1),
					MatchFlag:          true,
					MostSignificantBit: 32768,
					TagMaskNumBits:     32,
					TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
					TagDataNumBits:     32,
					TagData:            []byte{0x0, 0x1, 0x2, 0x3},
				},
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AccessSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 208, AccessSpecStopTrigger.
func TestAccessSpecStopTrigger_roundTrip(t *testing.T) {
	p := AccessSpecStopTrigger{
		Trigger:             AccessSpecStopTriggerOperationCount,
		OperationCountValue: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AccessSpecStopTrigger
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 209, AccessCommand.
func TestAccessCommand_roundTrip(t *testing.T) {
	p := AccessCommand{
		C1G2TagSpec: C1G2TagSpec{
			TagPattern1: C1G2TargetTag{
				C1G2MemoryBank:     uint8(1),
				MatchFlag:          true,
				MostSignificantBit: 32768,
				TagMaskNumBits:     32,
				TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
				TagDataNumBits:     32,
				TagData:            []byte{0x0, 0x1, 0x2, 0x3},
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AccessCommand
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 210, ClientRequestOpSpec.
func TestClientRequestOpSpec_roundTrip(t *testing.T) {
	p := ClientRequestOpSpec(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ClientRequestOpSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 211, ClientRequestResponse.
func TestClientRequestResponse_roundTrip(t *testing.T) {
	p := ClientRequestResponse{
		AccessSpecID: 2147483648,
		EPCData: EPCData{
			EPCNumBits: 32,
			EPC:        []byte{0x0, 0x1, 0x2, 0x3},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ClientRequestResponse
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 217, LLRPConfigurationStateValue.
func TestLLRPConfigurationStateValue_roundTrip(t *testing.T) {
	p := LLRPConfigurationStateValue(2147483648)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 LLRPConfigurationStateValue
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 218, Identification.
func TestIdentification_roundTrip(t *testing.T) {
	p := Identification{
		IDType:   ID_EPC,
		ReaderID: []byte{128, 1, 2, 3, 4, 5, 6, 7},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 Identification
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 219, GPOWriteData.
func TestGPOWriteData_roundTrip(t *testing.T) {
	p := GPOWriteData{
		Port: 1,
		Data: true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 GPOWriteData
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 220, KeepAliveSpec.
func TestKeepAliveSpec_roundTrip(t *testing.T) {
	p := KeepAliveSpec{
		Trigger:  KATriggerPeriodic,
		Interval: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 KeepAliveSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 221, AntennaProperties.
func TestAntennaProperties_roundTrip(t *testing.T) {
	p := AntennaProperties{
		AntennaConnected: true,
		AntennaID:        1,
		AntennaGain:      uint16(32768),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AntennaProperties
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 222, AntennaConfiguration.
func TestAntennaConfiguration_roundTrip(t *testing.T) {
	p := AntennaConfiguration{
		AntennaID: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AntennaConfiguration
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 223, RFReceiver.
func TestRFReceiver_roundTrip(t *testing.T) {
	p := RFReceiver(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFReceiver
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 224, RFTransmitter.
func TestRFTransmitter_roundTrip(t *testing.T) {
	p := RFTransmitter{
		HopTableID:         32768,
		ChannelIndex:       32768,
		TransmitPowerIndex: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFTransmitter
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 225, GPIPortCurrentState.
func TestGPIPortCurrentState_roundTrip(t *testing.T) {
	p := GPIPortCurrentState{
		Port:    1,
		Enabled: true,
		State:   GPIStateHigh,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 GPIPortCurrentState
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 226, EventsAndReports.
func TestEventsAndReports_roundTrip(t *testing.T) {
	p := EventsAndReports(true)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 EventsAndReports
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 237, ROReportSpec.
func TestROReportSpec_roundTrip(t *testing.T) {
	p := ROReportSpec{
		Trigger: NSecondsOrAIEnd,
		N:       32768,
		TagReportContentSelector: TagReportContentSelector{
			EnableROSpecID:             true,
			EnableSpecIndex:            true,
			EnableInventoryParamSpecID: true,
			EnableAntennaID:            true,
			EnableChannelIndex:         true,
			EnablePeakRSSI:             true,
			EnableFirstSeenTimestamp:   true,
			EnableLastSeenTimestamp:    true,
			EnableTagSeenCount:         true,
			EnableAccessSpecID:         true,
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROReportSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 238, TagReportContentSelector.
func TestTagReportContentSelector_roundTrip(t *testing.T) {
	p := TagReportContentSelector{
		EnableROSpecID:             true,
		EnableSpecIndex:            true,
		EnableInventoryParamSpecID: true,
		EnableAntennaID:            true,
		EnableChannelIndex:         true,
		EnablePeakRSSI:             true,
		EnableFirstSeenTimestamp:   true,
		EnableLastSeenTimestamp:    true,
		EnableTagSeenCount:         true,
		EnableAccessSpecID:         true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 TagReportContentSelector
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 239, AccessReportSpec.
func TestAccessReportSpec_roundTrip(t *testing.T) {
	p := AccessReportSpec(0)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AccessReportSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 240, TagReportData.
func TestTagReportData_roundTrip(t *testing.T) {
	p := TagReportData{
		EPCData: EPCData{
			EPCNumBits: 32,
			EPC:        []byte{0x0, 0x1, 0x2, 0x3},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 TagReportData
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 241, EPCData.
func TestEPCData_roundTrip(t *testing.T) {
	p := EPCData{
		EPCNumBits: 32,
		EPC:        []byte{0x0, 0x1, 0x2, 0x3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 EPCData
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 242, RFSurveyReportData.
func TestRFSurveyReportData_roundTrip(t *testing.T) {
	p := RFSurveyReportData{
		FrequencyRSSILevelEntries: []FrequencyRSSILevelEntry{
			{
				Frequency:    uint32(2147483648),
				Bandwidth:    uint32(2147483648),
				AverageRSSI:  int8(-1),
				PeakRSSI:     int8(-1),
				UTCTimestamp: UTCTimestamp(uint64(9223372036854775808)),
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFSurveyReportData
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 243, FrequencyRSSILevelEntry.
func TestFrequencyRSSILevelEntry_roundTrip(t *testing.T) {
	p := FrequencyRSSILevelEntry{
		Frequency:    uint32(2147483648),
		Bandwidth:    uint32(2147483648),
		AverageRSSI:  int8(-1),
		PeakRSSI:     int8(-1),
		UTCTimestamp: UTCTimestamp(uint64(9223372036854775808)),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FrequencyRSSILevelEntry
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 244, ReaderEventNotificationSpec.
func TestReaderEventNotificationSpec_roundTrip(t *testing.T) {
	p := ReaderEventNotificationSpec{
		EventNotificationStates: []EventNotificationState{
			{
				ReaderEventType:     NotifyRFSurvey,
				NotificationEnabled: true,
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ReaderEventNotificationSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 245, EventNotificationState.
func TestEventNotificationState_roundTrip(t *testing.T) {
	p := EventNotificationState{
		ReaderEventType:     NotifyRFSurvey,
		NotificationEnabled: true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 EventNotificationState
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 246, ReaderEventNotificationData.
func TestReaderEventNotificationData_roundTrip(t *testing.T) {
	p := ReaderEventNotificationData{
		UTCTimestamp: UTCTimestamp(uint64(9223372036854775808)),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ReaderEventNotificationData
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 247, HoppingEvent.
func TestHoppingEvent_roundTrip(t *testing.T) {
	p := HoppingEvent(32768)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 HoppingEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 248, GPIEvent.
func TestGPIEvent_roundTrip(t *testing.T) {
	p := GPIEvent{
		Port:  32768,
		Event: true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 GPIEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 249, ROSpecEvent.
func TestROSpecEvent_roundTrip(t *testing.T) {
	p := ROSpecEvent{
		Event:              ROSpecEnded,
		ROSpecID:           2147483648,
		PreemptingROSpecID: 2147483648,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ROSpecEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 250, ReportBufferLevelWarningEvent.
func TestReportBufferLevelWarningEvent_roundTrip(t *testing.T) {
	p := ReportBufferLevelWarningEvent(128)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ReportBufferLevelWarningEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 251, ReportBufferOverflowErrorEvent.
func TestReportBufferOverflowErrorEvent_roundTrip(t *testing.T) {
	p := ReportBufferOverflowErrorEvent{}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ReportBufferOverflowErrorEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 252, ReaderExceptionEvent.
func TestReaderExceptionEvent_roundTrip(t *testing.T) {
	p := ReaderExceptionEvent{
		Message: "some arbitrary string",
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ReaderExceptionEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 253, RFSurveyEvent.
func TestRFSurveyEvent_roundTrip(t *testing.T) {
	p := RFSurveyEvent{
		Event:    RFSurveyEnded,
		ROSpecID: 2147483648,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFSurveyEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 254, AISpecEvent.
func TestAISpecEvent_roundTrip(t *testing.T) {
	p := AISpecEvent{
		Event:     AISpecEnded,
		ROSpecID:  2147483648,
		SpecIndex: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AISpecEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 255, AntennaEvent.
func TestAntennaEvent_roundTrip(t *testing.T) {
	p := AntennaEvent{
		Event:     AntennaConnected,
		AntennaID: 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 AntennaEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 256, ConnectionAttemptEvent.
func TestConnectionAttemptEvent_roundTrip(t *testing.T) {
	p := ConnectionAttemptEvent(ConnExistsClientInitiated)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ConnectionAttemptEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 257, ConnectionCloseEvent.
func TestConnectionCloseEvent_roundTrip(t *testing.T) {
	p := ConnectionCloseEvent{}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ConnectionCloseEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 287, LLRPStatus.
func TestLLRPStatus_roundTrip(t *testing.T) {
	p := LLRPStatus{
		Status:           StatusMsgMsgUnexpected,
		ErrorDescription: "some arbitrary string",
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 LLRPStatus
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 288, FieldError.
func TestFieldError_roundTrip(t *testing.T) {
	p := FieldError{
		FieldIndex: 32768,
		ErrorCode:  StatusMsgMsgUnexpected,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 FieldError
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 289, ParameterError.
func TestParameterError_roundTrip(t *testing.T) {
	p := ParameterError{
		ParameterType: 0,
		ErrorCode:     StatusMsgMsgUnexpected,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 ParameterError
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 327, C1G2LLRPCapabilities.
func TestC1G2LLRPCapabilities_roundTrip(t *testing.T) {
	p := C1G2LLRPCapabilities{
		SupportsBlockErase:         true,
		SupportsBlockWrite:         true,
		SupportsBlockPermalock:     true,
		SupportsTagRecommissioning: true,
		SupportsUMIMethod2:         true,
		SupportsXPC:                true,
		MaxSelectFiltersPerQuery:   32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2LLRPCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 328, UHFC1G2RFModeTable.
func TestUHFC1G2RFModeTable_roundTrip(t *testing.T) {
	p := UHFC1G2RFModeTable{
		UHFC1G2RFModeTableEntries: []UHFC1G2RFModeTableEntry{
			{
				ModeID:                2147483648,
				DivideRatio:           DRSixtyFourToThree,
				IsEPCHagConformant:    true,
				Modulation:            Miller4,
				ForwardLinkModulation: SingleSidebandASK,
				SpectralMask:          SpectralMaskMultiInterrogator,
				BackscatterDataRate:   uint32(300000),
				PIERatio:              1500,
				MinTariTime:           uint32(9375),
				MaxTariTime:           uint32(9375),
				StepTariTime:          uint32(9375),
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 UHFC1G2RFModeTable
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 329, UHFC1G2RFModeTableEntry.
func TestUHFC1G2RFModeTableEntry_roundTrip(t *testing.T) {
	p := UHFC1G2RFModeTableEntry{
		ModeID:                2147483648,
		DivideRatio:           DRSixtyFourToThree,
		IsEPCHagConformant:    true,
		Modulation:            Miller4,
		ForwardLinkModulation: SingleSidebandASK,
		SpectralMask:          SpectralMaskMultiInterrogator,
		BackscatterDataRate:   uint32(300000),
		PIERatio:              1500,
		MinTariTime:           uint32(9375),
		MaxTariTime:           uint32(9375),
		StepTariTime:          uint32(9375),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 UHFC1G2RFModeTableEntry
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 330, C1G2InventoryCommand.
func TestC1G2InventoryCommand_roundTrip(t *testing.T) {
	p := C1G2InventoryCommand{
		TagInventoryStateAware: true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2InventoryCommand
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 331, C1G2Filter.
func TestC1G2Filter_roundTrip(t *testing.T) {
	p := C1G2Filter{
		TruncateAction: FilterActionDoNotTruncate,
		TagInventoryMask: C1G2TagInventoryMask{
			MemoryBank:         uint8(1),
			MostSignificantBit: 32768,
			TagMaskNumBits:     32,
			TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2Filter
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 332, C1G2TagInventoryMask.
func TestC1G2TagInventoryMask_roundTrip(t *testing.T) {
	p := C1G2TagInventoryMask{
		MemoryBank:         uint8(1),
		MostSignificantBit: 32768,
		TagMaskNumBits:     32,
		TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2TagInventoryMask
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 333, C1G2TagInventoryStateAwareFilterAction.
func TestC1G2TagInventoryStateAwareFilterAction_roundTrip(t *testing.T) {
	p := C1G2TagInventoryStateAwareFilterAction{
		Target:       InvTargetInventoriedS1,
		FilterAction: OnSelectMClearUSet,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2TagInventoryStateAwareFilterAction
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 334, C1G2TagInventoryStateUnawareFilterAction.
func TestC1G2TagInventoryStateUnawareFilterAction_roundTrip(t *testing.T) {
	p := C1G2TagInventoryStateUnawareFilterAction(0)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2TagInventoryStateUnawareFilterAction
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 335, C1G2RFControl.
func TestC1G2RFControl_roundTrip(t *testing.T) {
	p := C1G2RFControl{
		RFModeID: 32768,
		Tari:     uint16(32768),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2RFControl
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 336, C1G2SingulationControl.
func TestC1G2SingulationControl_roundTrip(t *testing.T) {
	p := C1G2SingulationControl{
		Session:        uint8(1),
		TagPopulation:  32768,
		TagTransitTime: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2SingulationControl
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 337, C1G2TagInventoryStateAwareSingulationAction.
func TestC1G2TagInventoryStateAwareSingulationAction_roundTrip(t *testing.T) {
	p := C1G2TagInventoryStateAwareSingulationAction{
		SessionState: SessionStateB,
		SLState:      SLStateDeasserted,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2TagInventoryStateAwareSingulationAction
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 338, C1G2TagSpec.
func TestC1G2TagSpec_roundTrip(t *testing.T) {
	p := C1G2TagSpec{
		TagPattern1: C1G2TargetTag{
			C1G2MemoryBank:     uint8(1),
			MatchFlag:          true,
			MostSignificantBit: 32768,
			TagMaskNumBits:     32,
			TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
			TagDataNumBits:     32,
			TagData:            []byte{0x0, 0x1, 0x2, 0x3},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2TagSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 339, C1G2TargetTag.
func TestC1G2TargetTag_roundTrip(t *testing.T) {
	p := C1G2TargetTag{
		C1G2MemoryBank:     uint8(1),
		MatchFlag:          true,
		MostSignificantBit: 32768,
		TagMaskNumBits:     32,
		TagMask:            []byte{0x0, 0x1, 0x2, 0x3},
		TagDataNumBits:     32,
		TagData:            []byte{0x0, 0x1, 0x2, 0x3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2TargetTag
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 341, C1G2Read.
func TestC1G2Read_roundTrip(t *testing.T) {
	p := C1G2Read{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2MemoryBank: uint8(1),
		WordAddress:    32768,
		WordCount:      32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2Read
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 342, C1G2Write.
func TestC1G2Write_roundTrip(t *testing.T) {
	p := C1G2Write{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2MemoryBank: uint8(1),
		WordAddress:    32768,
		Data:           []uint16{32768, 1, 2, 3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2Write
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 343, C1G2Kill.
func TestC1G2Kill_roundTrip(t *testing.T) {
	p := C1G2Kill{
		OpSpecID:     32768,
		KillPassword: 2147483648,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2Kill
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 344, C1G2Lock.
func TestC1G2Lock_roundTrip(t *testing.T) {
	p := C1G2Lock{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2LockPayloads: []C1G2LockPayload{
			{
				LockPrivilege: LockPrivPermaunlock,
				LockData:      LockDataEPCMemory,
			},
		},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2Lock
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 345, C1G2LockPayload.
func TestC1G2LockPayload_roundTrip(t *testing.T) {
	p := C1G2LockPayload{
		LockPrivilege: LockPrivPermaunlock,
		LockData:      LockDataEPCMemory,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2LockPayload
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 346, C1G2BlockErase.
func TestC1G2BlockErase_roundTrip(t *testing.T) {
	p := C1G2BlockErase{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2MemoryBank: uint8(1),
		WordAddress:    32768,
		WordCount:      32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2BlockErase
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 347, C1G2BlockWrite.
func TestC1G2BlockWrite_roundTrip(t *testing.T) {
	p := C1G2BlockWrite{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2MemoryBank: uint8(1),
		WordAddress:    32768,
		Data:           []uint16{32768, 1, 2, 3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2BlockWrite
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 348, C1G2EPCMemorySelector.
func TestC1G2EPCMemorySelector_roundTrip(t *testing.T) {
	p := C1G2EPCMemorySelector{
		CRCEnabled:     true,
		PCBitsEnabled:  true,
		XPCBitsEnabled: true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2EPCMemorySelector
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 349, C1G2ReadOpSpecResult.
func TestC1G2ReadOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2ReadOpSpecResult{
		C1G2ReadOpSpecResultType: 0,
		OpSpecID:                 32768,
		Data:                     []uint16{32768, 1, 2, 3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2ReadOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 350, C1G2WriteOpSpecResult.
func TestC1G2WriteOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2WriteOpSpecResult{
		C1G2WriteOpSpecResultType: 0,
		OpSpecID:                  32768,
		WordsWritten:              32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2WriteOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 351, C1G2KillOpSpecResult.
func TestC1G2KillOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2KillOpSpecResult{
		C1G2KillResult: 0,
		OpSpecID:       32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2KillOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 352, C1G2LockOpSpecResult.
func TestC1G2LockOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2LockOpSpecResult{
		C1G2LockResult: 0,
		OpSpecID:       32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2LockOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 353, C1G2BlockEraseOpSpecResult.
func TestC1G2BlockEraseOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2BlockEraseOpSpecResult{
		C1G2BlockEraseResult: 0,
		OpSpecID:             32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2BlockEraseOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 354, C1G2BlockWriteOpSpecResult.
func TestC1G2BlockWriteOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2BlockWriteOpSpecResult{
		C1G2BlockWriteResult: 0,
		OpSpecID:             32768,
		WordsWritten:         32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2BlockWriteOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 355, LoopSpec.
func TestLoopSpec_roundTrip(t *testing.T) {
	p := LoopSpec(2147483648)
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 LoopSpec
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 356, SpecLoopEvent.
func TestSpecLoopEvent_roundTrip(t *testing.T) {
	p := SpecLoopEvent{
		ROSpecID:  2147483648,
		LoopCount: 2147483648,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 SpecLoopEvent
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 357, C1G2Recommission.
func TestC1G2Recommission_roundTrip(t *testing.T) {
	p := C1G2Recommission{
		OpSpecID:     32768,
		KillPassword: 2147483648,
		SB3:          true,
		SB2:          true,
		LSB:          true,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2Recommission
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 358, C1G2BlockPermalock.
func TestC1G2BlockPermalock_roundTrip(t *testing.T) {
	p := C1G2BlockPermalock{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2MemoryBank: uint8(1),
		BlockAddress:   32768,
		BlockMask:      []uint16{32768, 1, 2, 3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2BlockPermalock
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 359, C1G2GetBlockPermalockStatus.
func TestC1G2GetBlockPermalockStatus_roundTrip(t *testing.T) {
	p := C1G2GetBlockPermalockStatus{
		OpSpecID:       32768,
		AccessPassword: 2147483648,
		C1G2MemoryBank: uint8(1),
		BlockAddress:   32768,
		BlockRange:     32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2GetBlockPermalockStatus
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 360, C1G2RecommissionOpSpecResult.
func TestC1G2RecommissionOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2RecommissionOpSpecResult{
		C1G2RecommissionResult: 0,
		OpSpecID:               32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2RecommissionOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 361, C1G2BlockPermalockOpSpecResult.
func TestC1G2BlockPermalockOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2BlockPermalockOpSpecResult{
		C1G2BlockPermalockResult: C1G2BPLockNoResponseFromTag,
		OpSpecID:                 32768,
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2BlockPermalockOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 362, C1G2GetBlockPermalockStatusOpSpecResult.
func TestC1G2GetBlockPermalockStatusOpSpecResult_roundTrip(t *testing.T) {
	p := C1G2GetBlockPermalockStatusOpSpecResult{
		C1G2GetBlockPermalockStatusResult: 0,
		OpSpecID:                          32768,
		PermalockStatuses:                 []uint16{32768, 1, 2, 3},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 C1G2GetBlockPermalockStatusOpSpecResult
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 363, MaximumReceiveSensitivity.
func TestMaximumReceiveSensitivity_roundTrip(t *testing.T) {
	p := MaximumReceiveSensitivity(int16(-1))
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 MaximumReceiveSensitivity
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 365, RFSurveyFrequencyCapabilities.
func TestRFSurveyFrequencyCapabilities_roundTrip(t *testing.T) {
	p := RFSurveyFrequencyCapabilities{
		MinFrequency: uint32(2147483648),
		MaxFrequency: uint32(2147483648),
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 RFSurveyFrequencyCapabilities
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}

// Test Parameter 1023, Custom.
func TestCustom_roundTrip(t *testing.T) {
	p := Custom{
		VendorID: 2147483648,
		Subtype:  2147483648,
		Data:     []byte{128, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
	}
	b, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	var p2 Custom
	if err := p2.UnmarshalBinary(b); err != nil {
		t.Errorf("%+v\n%# 02x\n%+v\n%+v", err, b, p, p2)
	}
	if !reflect.DeepEqual(p, p2) {
		t.Errorf("mismatch:\n%# 02x\n%+v\n%+v", b, p, p2)
	}
}
