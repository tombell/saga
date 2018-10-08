package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewGroupingField(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	grouping, err := field.NewGroupingField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewGroupingField err to be nil")
	}

	if grouping == nil {
		t.Fatal("expected grouping to not be nil")
	}
}

func TestNewGroupingFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E0074006900630073002000520072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewGroupingField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewGroupingField err to be ErrUnexpectedEOF")
	}
}

func TestNewGroupingFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000140000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewGroupingField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewGroupingField err to be ErrUnexpectedIdentifier")
	}
}

func TestGroupingValue(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	grouping, err := field.NewGroupingField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewGroupingField err to be nil")
	}

	actual := grouping.Value()
	expected := "Atlantics Records"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
