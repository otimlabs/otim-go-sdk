package signer

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/tkhq/go-sdk/pkg/api/models"
)

func ptr(s string) *string {
	return &s
}

func TestSigFromTurnkeyResult(t *testing.T) {
	tests := []struct {
		name    string
		input   *models.SignRawPayloadResult
		wantV   uint8
		wantR   string
		wantS   string
		wantErr bool
	}{
		{
			name:    "nil result",
			input:   nil,
			wantErr: true,
		},
		{
			name: "nil V",
			input: &models.SignRawPayloadResult{
				V: nil,
				R: ptr("abc"),
				S: ptr("def"),
			},
			wantErr: true,
		},
		{
			name: "nil R",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: nil,
				S: ptr("def"),
			},
			wantErr: true,
		},
		{
			name: "nil S",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("abc"),
				S: nil,
			},
			wantErr: true,
		},
		{
			name: "valid",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("abcd1234"),
				S: ptr("5678efab"),
			},
			wantV: 27,
			wantR: "0xabcd1234",
			wantS: "0x5678efab",
		},
		{
			name: "leading zeros in R (bug case)",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("00abcd1234"),
				S: ptr("5678efab"),
			},
			wantV: 27,
			wantR: "0xabcd1234",
			wantS: "0x5678efab",
		},
		{
			name: "leading zeros in S (bug case)",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("abcd1234"),
				S: ptr("005678efab"),
			},
			wantV: 27,
			wantR: "0xabcd1234",
			wantS: "0x5678efab",
		},
		{
			name: "leading zeros in both R and S",
			input: &models.SignRawPayloadResult{
				V: ptr("28"),
				R: ptr("00abcd1234"),
				S: ptr("005678efab"),
			},
			wantV: 28,
			wantR: "0xabcd1234",
			wantS: "0x5678efab",
		},
		{
			name: "multiple leading zeros",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("0000abcd1234"),
				S: ptr("0000005678efab"),
			},
			wantV: 27,
			wantR: "0xabcd1234",
			wantS: "0x5678efab",
		},
		{
			name: "full 32-byte signature components",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("00c9d5c5a3e8f1b2d4e6a8c0e2f4a6b8d0e2f4a6b8d0e2f4a6b8d0e2f4a6b8d0"),
				S: ptr("00a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1"),
			},
			wantV: 27,
			wantR: "0xc9d5c5a3e8f1b2d4e6a8c0e2f4a6b8d0e2f4a6b8d0e2f4a6b8d0e2f4a6b8d0",
			wantS: "0xa1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1",
		},
		{
			name: "invalid V (not a number)",
			input: &models.SignRawPayloadResult{
				V: ptr("abc"),
				R: ptr("abcd1234"),
				S: ptr("5678efab"),
			},
			wantErr: true,
		},
		{
			name: "invalid R (not hex)",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("ghij"),
				S: ptr("5678efab"),
			},
			wantErr: true,
		},
		{
			name: "invalid S (not hex)",
			input: &models.SignRawPayloadResult{
				V: ptr("27"),
				R: ptr("abcd1234"),
				S: ptr("ghij"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig, err := SigFromTurnkeyResult(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if sig.V != tt.wantV {
				t.Errorf("V = %d, want %d", sig.V, tt.wantV)
			}

			wantR, _ := uint256.FromHex(tt.wantR)
			if !sig.R.Eq(wantR) {
				t.Errorf("R = %s, want %s", sig.R.Hex(), wantR.Hex())
			}

			wantS, _ := uint256.FromHex(tt.wantS)
			if !sig.S.Eq(wantS) {
				t.Errorf("S = %s, want %s", sig.S.Hex(), wantS.Hex())
			}
		})
	}
}

func TestSigsFromTurnkeyBatchResult(t *testing.T) {
	tests := []struct {
		name    string
		input   []*models.SignRawPayloadResult
		wantLen int
		wantErr bool
	}{
		{
			name:    "nil results",
			input:   nil,
			wantErr: true,
		},
		{
			name:    "empty results",
			input:   []*models.SignRawPayloadResult{},
			wantLen: 0,
		},
		{
			name: "single result with leading zeros",
			input: []*models.SignRawPayloadResult{
				{
					V: ptr("27"),
					R: ptr("00abcd1234"),
					S: ptr("005678efab"),
				},
			},
			wantLen: 1,
		},
		{
			name: "multiple results",
			input: []*models.SignRawPayloadResult{
				{
					V: ptr("27"),
					R: ptr("00abcd1234"),
					S: ptr("5678efab"),
				},
				{
					V: ptr("28"),
					R: ptr("deadbeef"),
					S: ptr("00cafebabe"),
				},
			},
			wantLen: 2,
		},
		{
			name: "error in one result",
			input: []*models.SignRawPayloadResult{
				{
					V: ptr("27"),
					R: ptr("abcd1234"),
					S: ptr("5678efab"),
				},
				{
					V: ptr("invalid"),
					R: ptr("deadbeef"),
					S: ptr("cafebabe"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sigs, err := SigsFromTurnkeyBatchResult(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(sigs) != tt.wantLen {
				t.Errorf("got %d signatures, want %d", len(sigs), tt.wantLen)
			}
		})
	}
}