package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewField39Field(t *testing.T) {
	data, _ := hex.DecodeString("000000270000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	field, err := field.NewField39Field(hdr, buf)
	if err != nil {
		t.Fatal("expected NewField39Field err to be nil")
	}

	if field == nil {
		t.Fatal("expected field to not be nil")
	}
}

func TestNewField39FieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000002700000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewField39Field(hdr, buf)
	if err != io.EOF {
		t.Fatal("expected NewField39Field err to be EOF")
	}
}

func TestNewField39FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000490000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewField39Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewField39Field err to be ErrUnexpectedIdentifier")
	}
}

func TestField39Value(t *testing.T) {
	data, _ := hex.DecodeString("000000270000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	field, err := field.NewField39Field(hdr, buf)
	if err != nil {
		t.Fatal("expected NewField39Field err to be nil")
	}

	actual := field.Value()
	expected := []byte{1}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
