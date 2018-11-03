package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewCommentField(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	comment, err := field.NewCommentField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewCommentField err to be nil")
	}

	if comment == nil {
		t.Fatal("expected comment to not be nil")
	}
}

func TestNewCommentFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F0072007900")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewCommentField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewCommentField err to be ErrUnexpectedEOF")
	}
}

func TestNewCommentFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000120000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewCommentField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewCommentField err to be ErrUnexpectedIdentifier")
	}
}

func TestCommentValue(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	comment, err := field.NewCommentField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewCommentField err to be nil")
	}

	actual := comment.Value()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
