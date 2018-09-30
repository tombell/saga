package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewField69Field(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	field, err := field.NewField69Field(hdr, buf)
	if err != nil {
		t.Error("expected NewField69Field err to be nil")
	}

	if field == nil {
		t.Error("expected field to not be nil")
	}
}

func TestNewField69FieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewField69Field(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewField69Field err to be ErrUnexpectedEOF")
	}
}

func TestNewField69FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000440000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewField69Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Error("expected NewField69Field err to be ErrUnexpectedIdentifier")
	}
}

func TestField69Value(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	field, err := field.NewField69Field(hdr, buf)
	if err != nil {
		t.Error("expected NewField69Field err to be nil")
	}

	actual := field.Value()
	expected := []byte{0, 0, 0, 0}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
