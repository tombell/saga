package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewArtistField(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewArtistField err to be nil")
	}

	if artist == nil {
		t.Fatal("expected artist to not be nil")
	}
}

func TestNewArtistFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A002000460061007600")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewArtistField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewArtistField err to be ErrUnexpectedEOF")
	}
}

func TestNewArtistFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000005000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewArtistField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewArtistField err to be ErrUnexpectedIdentifier")
	}
}

func TestArtistValue(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewArtistField err to be nil")
	}

	actual := artist.Value()
	expected := "DJ Favorite, DJ Kharitonov"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
