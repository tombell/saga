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
		t.Fatal("expected NewHeader err to be nil")
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewAddedField err to be nil")
	}

	if added == nil {
		t.Fatal("expected added to not be nil")
	}
}

func TestNewAddedFieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewAddedField(hdr, buf)
	if err != io.EOF {
		t.Fatal("expected NewAddedField err to be EOF")
	}
}

func TestNewAddedFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000330000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewAddedField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewAddedField err to be ErrUnexpectedIdentifier")
	}
}

func TestAddedValue(t *testing.T) {
	data, _ := hex.DecodeString("000000340000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewAddedField err to be nil")
	}

	actual := added.Value()
	expected := false

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
