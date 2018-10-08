package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewGenreField(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	genre, err := field.NewGenreField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewGenreField err to be nil")
	}

	if genre == nil {
		t.Fatal("expected genre to not be nil")
	}
}

func TestNewGenreFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewGenreField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewGenreField err to be ErrUnexpectedEOF")
	}
}

func TestNewGenreFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewGenreField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewGenreField err to be ErrUnexpectedIdentifier")
	}
}

func TestGenreValue(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	genre, err := field.NewGenreField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewGenreField err to be nil")
	}

	actual := genre.Value()
	expected := "House"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
