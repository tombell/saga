package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewRemixerField(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	remixer, err := field.NewRemixerField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewRemixerField err to be nil")
	}

	if remixer == nil {
		t.Fatal("expected remixer to not be nil")
	}
}

func TestNewRemixerFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C00650073")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewRemixerField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewRemixerField err to be ErrUnexpectedEOF")
	}
}

func TestNewRemixerFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000001500000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewRemixerField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewRemixerField err to be ErrUnexpectedIdentifier")
	}
}

func TestRemixerValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	remixer, err := field.NewRemixerField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewRemixerField err to be nil")
	}

	actual := remixer.Value()
	expected := "Frankie Knuckles"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
