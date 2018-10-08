package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func generateBytes(data string) []byte {
	b, _ := hex.DecodeString(data)
	return b
}

func TestNewHeader(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader to return nil error")
	}

	if hdr.Identifier != 1 {
		t.Fatal("expected identifier to be 1")
	}

	if hdr.Length != 4 {
		t.Fatal("expected length to be 4")
	}
}

func TestNewHeaderUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000102")
	buf := bytes.NewBuffer(data)

	_, err := field.NewHeader(buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected err to be unexpected eof error")
	}
}

func TestHeaderString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"artist", generateBytes("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000"), "Field: 7, Data length: 54"},
		{"title", generateBytes("00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000"), "Field: 6, Data length: 68"},
		{"starttime", generateBytes("0000001C000000045B903D08"), "Field: 28, Data length: 4"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.input)

			hdr, err := field.NewHeader(buf)
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
