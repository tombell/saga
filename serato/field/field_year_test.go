package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewYearField(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A00320030003100380000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	year, err := field.NewYearField(hdr, buf)
	if err != nil {
		t.Error("expected NewYearField err to be nil")
	}

	if year == nil {
		t.Error("expected year to not be nil")
	}
}

func TestNewYearFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A0032003000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewYearField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewYearField err to be ErrUnexpectedEOF")
	}
}

func TestYearValue(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A00320030003100380000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	year, err := field.NewYearField(hdr, buf)
	if err != nil {
		t.Error("expected NewYearField err to be nil")
	}

	actual := year.Value()
	expected := "2018"

	if actual != expected {
		t.Errorf("expected value to be %s, got %s", expected, actual)
	}
}
