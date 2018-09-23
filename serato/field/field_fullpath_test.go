package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/field"
)

func TestNewFullPathField(t *testing.T) {
	data, _ := hex.DecodeString("00000002000000CA002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D007000330000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	fullpath, err := field.NewFullPathField(hdr, buf)
	if err != nil {
		t.Error("expected NewFullPathField err to be nil")
	}

	if fullpath == nil {
		t.Error("expected fullpath to not be nil")
	}
}

func TestNewFullPathFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000002000000CA002F00550073006500720073002F0074006F006D006")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	_, err = field.NewFullPathField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Error("expected NewFullPathField err to be ErrUnexpectedEOF")
	}
}

func TestFullPathValue(t *testing.T) {
	data, _ := hex.DecodeString("00000002000000CA002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D007000330000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Error("expected NewHeader err to be nil")
	}

	fullpath, err := field.NewFullPathField(hdr, buf)
	if err != nil {
		t.Error("expected NewFullPathField err to be nil")
	}

	actual := fullpath.Value()
	expected := "/Users/tombell/Music/__ New __/Classic House Summer '18/10947360_Do_You_Wanna_House_Original_Mix.mp3"

	if actual != expected {
		t.Errorf("expected value to be %v, got %v", expected, actual)
	}
}
