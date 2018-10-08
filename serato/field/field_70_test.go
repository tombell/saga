package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewField70Field(t *testing.T) {
	data, _ := hex.DecodeString("000000460000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	field, err := field.NewField70Field(hdr, buf)
	if err != nil {
		t.Fatal("expected NewField70Field err to be nil")
	}

	if field == nil {
		t.Fatal("expected field to not be nil")
	}
}

func TestNewField70FieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000004600000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewField70Field(hdr, buf)
	if err != io.EOF {
		t.Fatal("expected NewField70Field err to be EOF")
	}
}

func TestNewField70FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000440000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewField70Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewField70Field err to be ErrUnexpectedIdentifier")
	}
}

func TestField70Value(t *testing.T) {
	data, _ := hex.DecodeString("000000460000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	field, err := field.NewField70Field(hdr, buf)
	if err != nil {
		t.Fatal("expected NewField70Field err to be nil")
	}

	actual := field.Value()
	expected := byte(0)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
