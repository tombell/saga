package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/chunk"
)

func generateBytes(data string) []byte {
	b, _ := hex.DecodeString(data)
	return b
}

func TestNewHeader(t *testing.T) {
	data, _ := hex.DecodeString("7672736E0000003C")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader to return nil error")
	}

	if hdr.Identifier != [4]byte{0x76, 0x72, 0x73, 0x6E} {
		t.Fatal("expected identifier to be 0x7672736E")
	}

	if hdr.Length != 60 {
		t.Fatal("expected length to be 60")
	}
}

func TestNewHeaderUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000102")
	buf := bytes.NewBuffer(data)

	_, err := chunk.NewHeader(buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected err to be unexpected eof error")
	}
}

func TestHeaderType(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		expectedType string
	}{
		{"vrsn", generateBytes("7672736E0000003C"), "vrsn"},
		{"oent", generateBytes("6F656E740000028F"), "oent"},
		{"adat", generateBytes("6164617400000287"), "adat"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.input)

			hdr, err := chunk.NewHeader(buf)
			if err != nil {
				t.Fatal("expected err to be nil")
			}

			actual := hdr.Type()
			if actual != tc.expectedType {
				t.Fatalf("expected type to be %s, got %s", tc.expectedType, actual)
			}
		})
	}
}

func TestHeaderString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"vrsn", generateBytes("7672736E0000003C"), "Chunk: vrsn, Data length: 60"},
		{"oent", generateBytes("6F656E740000028F"), "Chunk: oent, Data length: 655"},
		{"adat", generateBytes("6164617400000287"), "Chunk: adat, Data length: 647"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.input)

			hdr, err := chunk.NewHeader(buf)
			if err != nil {
				t.Fatal("expected err to be nil")
			}

			actual := hdr.String()
			if actual != tc.expected {
				t.Fatalf("expected string to be %s, got %s", tc.expected, actual)
			}
		})
	}
}
