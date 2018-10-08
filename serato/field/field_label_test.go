package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewLabelField(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	label, err := field.NewLabelField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewLabelField err to be nil")
	}

	if label == nil {
		t.Fatal("expected label to not be nil")
	}
}

func TestNewLabelFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewLabelField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewLabelField err to be ErrUnexpectedEOF")
	}
}

func TestNewLabelFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000140000002400410074006C0061006E00740069006300730020005200650063006F0000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewLabelField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewLabelField err to be ErrUnexpectedIdentifier")
	}
}

func TestLabelValue(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	label, err := field.NewLabelField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewLabelField err to be nil")
	}

	actual := label.Value()
	expected := "Atlantics Records"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
