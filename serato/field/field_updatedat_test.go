package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/tombell/saga/serato/field"
)

func TestNewUpdatedAtField(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	updatedat, err := field.NewUpdatedAtField(hdr, buf)
	if err != nil {
		t.Error("expected NewUpdatedAtField err to be nil")
	}

	if updatedat == nil {
		t.Error("expected updatedat to not be nil")
	}
}

func TestNewUpdatedAtFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewUpdatedAtField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewUpdatedAtField err to be ErrUnexpectedEOF")
	}
}

func TestNewUpdatedAtFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000045000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewUpdatedAtField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Error("expected NewUpdatedAtField err to be ErrUnexpectedIdentifier")
	}
}

func TestUpdatedAtValue(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	updatedat, err := field.NewUpdatedAtField(hdr, buf)
	if err != nil {
		t.Error("expected NewUpdatedAtField err to be nil")
	}

	actual := updatedat.Value()
	expected := time.Date(2018, time.September, 5, 20, 33, 39, 0, time.UTC)

	if actual != expected {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
