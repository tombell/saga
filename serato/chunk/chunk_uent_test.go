package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/chunk"
)

func TestNewUentChunk(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewUentCHunk err to be nil")
	}

	if uent == nil {
		t.Fatal("expected uent to not be nil")
	}
}

func TestNewUentChunkUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("75656E7400000004000000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewUentChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewUentCHunk err to be ErrUnexpectedEOF")
	}
}

func TestNewUentChunkUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("74656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewUentChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatal("expected NewUentCHunk err to be ErrUnexpectedEOF")
	}
}

func TestUentHeader(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewUentCHunk err to be nil")
	}

	if uent.Header() != hdr {
		t.Fatal("expected header to be the same")
	}
}

func TestUentType(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewUentChunk err to be nil")
	}

	actual := uent.Type()
	expected := "uent"

	if actual != expected {
		t.Fatalf("expected type to be %s, got %s", expected, actual)
	}
}
