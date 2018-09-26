package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewAddedField(t *testing.T) {
	data, _ := hex.DecodeString("000000340000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Error("expected NewAddedField err to be nil")
	}

	if added == nil {
		t.Error("expected added to not be nil")
	}
}

func TestNewAddedFieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewAddedField(hdr, buf)
	if err != io.EOF {
		t.Error("expected NewAddedField err to be EOF")
	}
}

func TestAddedValue(t *testing.T) {
	data, _ := hex.DecodeString("000000340000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Error("expected NewAddedField err to be nil")
	}

	actual := added.Value()
	expected := byte(0)

	if actual != expected {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
