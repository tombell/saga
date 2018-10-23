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
		t.Fatal("expected NewHeader err to be nil")
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewPlayedField err to be nil")
	}

	if played == nil {
		t.Fatal("expected played to not be nil")
	}
}

func TestNewPlayedFieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003200000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewPlayedField(hdr, buf)
	if err != io.EOF {
		t.Fatal("expected NewPlayedField err to be EOF")
	}
}

func TestNewPlayedFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000330000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewPlayedField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewPlayedField err to be ErrUnexpectedIdentifier")
	}
}

func TestPlayedValue(t *testing.T) {
	data, _ := hex.DecodeString("000000320000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewPlayedField err to be nil")
	}

	actual := played.Value()
	expected := true

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
