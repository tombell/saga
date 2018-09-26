package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/tombell/saga/serato/field"
)

func TestNewEndTimeField(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	endtime, err := field.NewEndTimeField(hdr, buf)
	if err != nil {
		t.Error("expected NewEndTimeField err to be nil")
	}

	if endtime == nil {
		t.Error("expected endtime to not be nil")
	}
}

func TestNewEndTimeFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045BDA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewEndTimeField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewEndTimeField err to be ErrUnexpectedEOF")
	}
}

func TestEndTimeValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	endtime, err := field.NewEndTimeField(hdr, buf)
	if err != nil {
		t.Error("expected NewEndTimeField err to be nil")
	}

	actual := endtime.Value()
	expected := time.Date(2018, time.September, 5, 20, 33, 39, 0, time.UTC)

	if actual != expected {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
