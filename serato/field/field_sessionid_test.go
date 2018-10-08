package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewSessionIDField(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	sessionid, err := field.NewSessionIDField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewSessionIDField err to be nil")
	}

	if sessionid == nil {
		t.Fatal("expected sessionid to not be nil")
	}
}

func TestNewSessionIDFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000002")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewSessionIDField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewSessionIDField err to be ErrUnexpectedEOF")
	}
}

func TestNewSessionIDFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000004000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewSessionIDField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewSessionIDField err to be ErrUnexpectedIdentifier")
	}
}

func TestSessionIDValue(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	sessionid, err := field.NewSessionIDField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewSessionIDField err to be nil")
	}

	actual := sessionid.Value()
	expected := 210

	if actual != expected {
		t.Fatalf("expected value to be %d, got %d", expected, actual)
	}
}
