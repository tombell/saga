package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewPlayedField(t *testing.T) {
	data, _ := hex.DecodeString("000000320000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Error("expected NewPlayedField err to be nil")
	}

	if played == nil {
		t.Error("expected played to not be nil")
	}
}

func TestNewPlayedFieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003200000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewPlayedField(hdr, buf)
	if err != io.EOF {
		t.Error("expected NewPlayedField err to be EOF")
	}
}

func TestPlayedValue(t *testing.T) {
	data, _ := hex.DecodeString("000000320000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Error("expected NewPlayedField err to be nil")
	}

	actual := played.Value()
	expected := byte(1)

	if actual != expected {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
