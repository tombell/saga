package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewDeckField(t *testing.T) {
	data, _ := hex.DecodeString("0000001F0000000400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	deck, err := field.NewDeckField(hdr, buf)
	if err != nil {
		t.Error("expected NewDeckField err to be nil")
	}

	if deck == nil {
		t.Error("expected deck to not be nil")
	}
}

func TestNewDeckFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001F00000004000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewDeckField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewDeckField err to be ErrUnexpectedEOF")
	}
}

func TestDeckValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001F0000000400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	deck, err := field.NewDeckField(hdr, buf)
	if err != nil {
		t.Error("expected NewDeckField err to be nil")
	}

	actual := deck.Value()
	expected := 1

	if actual != expected {
		t.Errorf("expected value to be %d, got %d", expected, actual)
	}
}
