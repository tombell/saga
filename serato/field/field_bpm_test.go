package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewBPMField(t *testing.T) {
	data, _ := hex.DecodeString("0000000F0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	bpm, err := field.NewBPMField(hdr, buf)
	if err != nil {
		t.Error("expected NewBPMField err to be nil")
	}

	if bpm == nil {
		t.Error("expected bpm to not be nil")
	}
}

func TestNewBPMFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000000F000000040000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewBPMField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewBPMField err to be ErrUnexpectedEOF")
	}
}

func TestNewBPMFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000D0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewBPMField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Error("expected NewBPMField err to be ErrUnexpectedIdentifier")
	}
}

func TestBPMValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000F0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	bpm, err := field.NewBPMField(hdr, buf)
	if err != nil {
		t.Error("expected NewBPMField err to be nil")
	}

	actual := bpm.Value()
	expected := 119

	if actual != expected {
		t.Errorf("expected value to be %d, got %d", expected, actual)
	}
}
