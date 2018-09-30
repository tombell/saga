package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewPlayTimeField(t *testing.T) {
	data, _ := hex.DecodeString("0000002D000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	playtime, err := field.NewPlayTimeField(hdr, buf)
	if err != nil {
		t.Error("expected NewPlayTimeField err to be nil")
	}

	if playtime == nil {
		t.Error("expected playtime to not be nil")
	}
}

func TestNewPlayTimeFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000002D0000000400009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewPlayTimeField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewPlayTimeField err to be ErrUnexpectedEOF")
	}
}

func TestNewPlayTimeFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000002E000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewPlayTimeField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Error("expected NewPlayTimeField err to be ErrUnexpectedIdentifier")
	}
}

func TestPlayTimeValue(t *testing.T) {
	data, _ := hex.DecodeString("0000002D000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	playtime, err := field.NewPlayTimeField(hdr, buf)
	if err != nil {
		t.Error("expected NewPlayTimeField err to be nil")
	}

	actual := playtime.Value()
	expected := 155

	if actual != expected {
		t.Errorf("expected value to be %d, got %d", expected, actual)
	}
}
