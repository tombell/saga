package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/chunk"
)

func TestNewHeader(t *testing.T) {
	data, _ := hex.DecodeString("7672736E0000003C")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)

	if err != nil {
		t.Fatal("expected NewHeader to return nil error")
	}

	if hdr.Identifier != [4]byte{0x76, 0x72, 0x73, 0x6E} {
		t.Error("expected identifier to be 0x7672736E")
	}

	if hdr.Length != 60 {
		t.Error("expected length to be 60")
	}

	if hdr.Type() != "vrsn" {
		t.Error("expected type to be vrsn")
	}

}

func TestNewHeaderError(t *testing.T) {
	data := []byte{00, 01, 02}
	buf := bytes.NewBuffer(data)

	_, err := chunk.NewHeader(buf)

	if err != io.ErrUnexpectedEOF {
		t.Error("expected err to be unexpected eof error")
	}
}
