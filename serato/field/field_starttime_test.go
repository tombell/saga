package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/tombell/saga/serato/field"
)

func TestNewStartTimeField(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	starttime, err := field.NewStartTimeField(hdr, buf)
	if err != nil {
		t.Error("expected NewStartTimeField err to be nil")
	}

	if starttime == nil {
		t.Error("expected starttime to not be nil")
	}
}

func TestNewStartTimeFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B908")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewStartTimeField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewStartTimeField err to be ErrUnexpectedEOF")
	}
}

func TestNewStartTimeFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000002C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewStartTimeField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Error("expected NewStartTimeField err to be ErrUnexpectedIdentifier")
	}
}

func TestStartTimeValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	starttime, err := field.NewStartTimeField(hdr, buf)
	if err != nil {
		t.Error("expected NewStartTimeField err to be nil")
	}

	actual := starttime.Value()
	expected := time.Date(2018, time.September, 5, 20, 31, 04, 0, time.UTC)

	if actual != expected {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
