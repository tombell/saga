package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/chunk"
)

func TestNewVrsnChunk(t *testing.T) {
	data, _ := hex.DecodeString("7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Error("expected NewVrsnChunk err to be nil")
	}

	if vrsn == nil {
		t.Error("expected vrsn to not be nil")
	}
}

func TestNewVrsnChunkUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("7672736E000000376004500200052006500760067")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = chunk.NewVrsnChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewVrsnChunk err to be ErrUnexpectedEOF")
	}
}

func TestVrsnHeader(t *testing.T) {
	data, _ := hex.DecodeString("7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Error("expected NewVrsnChunk err to be nil")
	}

	if vrsn.Header() != hdr {
		t.Error("expected header to be the same")
	}
}

func TestVrsnType(t *testing.T) {
	data, _ := hex.DecodeString("7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Error("expected NewVrsnChunk err to be nil")
	}

	actual := vrsn.Type()
	expected := "vrsn"

	if actual != expected {
		t.Errorf("expected type to be %s, got %s", expected, actual)
	}
}

func TestVrsnVersion(t *testing.T) {
	data, _ := hex.DecodeString("7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Error("expected NewVrsnChunk err to be nil")
	}

	actual := vrsn.Version()
	expected := "1.0/Serato Scratch LIVE Review"

	if actual != expected {
		t.Errorf("expected version to be %s, got %s", expected, actual)
	}
}
