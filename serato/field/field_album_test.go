package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewAlbumField(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	album, err := field.NewAlbumField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewAlbumField err to be nil")
	}

	if album == nil {
		t.Fatal("expected album to not be nil")
	}
}

func TestNewAlbumFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F00720079000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewAlbumField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewAlbumField err to be ErrUnexpectedEOF")
	}
}

func TestNewAlbumFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewAlbumField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewAlbumField err to be ErrUnexpectedIdentifier")
	}
}

func TestAlbumValue(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	album, err := field.NewAlbumField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewAlbumField err to be nil")
	}

	actual := album.Value()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
