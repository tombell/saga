package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/chunk"
)

func TestNewOrenChunk(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	oren, err := chunk.NewOrenChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewOrenChunk err to be nil")
	}

	if oren == nil {
		t.Fatal("expected oren to not be nil")
	}
}

func TestNewOrenChunkUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewOrenChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewOrenChunk err to be ErrUnexpectedEOF")
	}
}

func TestNewOrenChunkUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("6E72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewOrenChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatal("expected NewOrenChunk err to be ErrUnexpectedIdentifier")
	}
}

func TestOrenHeader(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	oren, err := chunk.NewOrenChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewOrenChunk err to be nil")
	}

	if oren.Header() != hdr {
		t.Fatal("expected header to be the same")
	}
}

func TestOrenType(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	oren, err := chunk.NewOrenChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewOrenChunk err to be nil")
	}

	actual := oren.Type()
	expected := "oren"

	if actual != expected {
		t.Fatalf("expected type to be %s, got %s", expected, actual)
	}
}
