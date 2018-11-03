package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewComposerField(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	composer, err := field.NewComposerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewComposerField err to be nil (%v)", err)
	}

	if composer == nil {
		t.Fatal("expected composer to not be nil")
	}
}

func TestNewComposerFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F0072007900")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewComposerField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewComposerField err to be ErrUnexpectedEOF")
	}
}

func TestNewComposerFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000150000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewComposerField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewComposerField err to be ErrUnexpectedIdentifier")
	}
}

func TestComposerValue(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	composer, err := field.NewComposerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewComposerField err to be nil (%v)", err)
	}

	actual := composer.Value()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
