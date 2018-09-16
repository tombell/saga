package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewHeader(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader to return nil error")
	}

	if hdr.Identifier != 1 {
		t.Error("expected identifier to be 1")
	}

	if hdr.Length != 4 {
		t.Error("expected length to be 4")
	}
}

func TestNewHeaderUnexpectedEOF(t *testing.T) {
	data := []byte{00, 01, 02}
	buf := bytes.NewBuffer(data)

	_, err := field.NewHeader(buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected err to be unexpected eof error")
	}
}
