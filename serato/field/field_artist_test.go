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
		t.Error("expected NewHeader err to be nil")
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Error("expected NewArtistField err to be nil")
	}

	if artist == nil {
		t.Error("expected artist to not be nil")
	}
}

func TestNewArtistFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F007")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewArtistField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewArtistField err to be ErrUnexpectedEOF")
	}
}

func TestArtistValue(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Error("expected NewArtistField err to be nil")
	}

	actual := artist.Value()
	expected := "DJ Favorite, DJ Kharitonov"

	if actual != expected {
		t.Errorf("expected value to be %s, got %s", expected, actual)
	}
}
