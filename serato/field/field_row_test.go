package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewRowField(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	row, err := field.NewRowField(hdr, buf)
	if err != nil {
		t.Error("expected NewRowField err to be nil")
	}

	if row == nil {
		t.Error("expected row to not be nil")
	}
}

func TestNewRowFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000010000000400000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewRowField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewRowField err to be ErrUnexpectedEOF")
	}
}

func TestNewRowFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000200000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewRowField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Error("expected NewRowField err to be ErrUnexpectedIdentifier")
	}
}

func TestRowValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	row, err := field.NewRowField(hdr, buf)
	if err != nil {
		t.Error("expected NewRowField err to be nil")
	}

	actual := row.Value()
	expected := 212

	if actual != expected {
		t.Errorf("expected value to be %d, got %d", expected, actual)
	}
}
