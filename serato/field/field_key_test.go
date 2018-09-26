package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewKeyField(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006D0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	key, err := field.NewKeyField(hdr, buf)
	if err != nil {
		t.Error("expected NewKeyField err to be nil")
	}

	if key == nil {
		t.Error("expected key to not be nil")
	}
}

func TestNewKeyFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewKeyField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewKeyField err to be ErrUnexpectedEOF")
	}
}

func TestKeyValue(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006D0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	key, err := field.NewKeyField(hdr, buf)
	if err != nil {
		t.Error("expected NewKeyField err to be nil")
	}

	actual := key.Value()
	expected := "Cm"

	if actual != expected {
		t.Errorf("expected value to be %s, got %s", expected, actual)
	}
}
