package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewTitleField(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	title, err := field.NewTitleField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewTitleField err to be nil")
	}

	if title == nil {
		t.Fatal("expected title to not be nil")
	}
}

func TestNewTitleFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F007500200")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewTitleField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewTitleField err to be ErrUnexpectedEOF")
	}
}

func TestNewTitleFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewTitleField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewTitleField err to be ErrUnexpectedIdentifier")
	}
}

func TestTitleValue(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	title, err := field.NewTitleField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewTitleField err to be nil")
	}

	actual := title.Value()
	expected := "Do You Wanna House (Original Mix)"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
