package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/saga/serato/chunk"
)

func TestNewOentChunk(t *testing.T) {
	data, _ := hex.DecodeString("6F656E740000028F61646174000002870000000100000004000000D400000002000000CA002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D00700033000000000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D006900780029000000000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000000000090000000C0048006F00750073006500000000000F0000000400000077000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000170000000A003200300031003800000000001C000000045B903D080000001D000000045B903DA30000001F00000004000000010000002D000000040000009B0000003000000004000000D200000032000000010100000033000000060043006D000000000034000000010000000035000000045B903DA3000000440000000400000000000000450000000400000000000000460000000100000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	oent, err := chunk.NewOentChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewOentChunk err to be nil")
	}

	if oent == nil {
		t.Fatal("expected oent to not be nil")
	}
}

func TestNewOentChunkUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("6F656E740000028F61646174000002870000000100000004000000D400000002000000CA002F0055007300650072007300")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewOentChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewOentChunk err to be ErrUnexpectedEOF")
	}
}

func TestNewOentChunkUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("6E656E740000028F61646174000002870000000100000004000000D400000002000000CA002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D00700033000000000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D006900780029000000000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000000000090000000C0048006F00750073006500000000000F0000000400000077000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000170000000A003200300031003800000000001C000000045B903D080000001D000000045B903DA30000001F00000004000000010000002D000000040000009B0000003000000004000000D200000032000000010100000033000000060043006D000000000034000000010000000035000000045B903DA3000000440000000400000000000000450000000400000000000000460000000100000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewOentChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatal("expected NewOentChunk err to be ErrUnexpectedIdentifier")
	}
}

func TestOentHeader(t *testing.T) {
	data, _ := hex.DecodeString("6F656E740000028F61646174000002870000000100000004000000D400000002000000CA002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D00700033000000000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D006900780029000000000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000000000090000000C0048006F00750073006500000000000F0000000400000077000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000170000000A003200300031003800000000001C000000045B903D080000001D000000045B903DA30000001F00000004000000010000002D000000040000009B0000003000000004000000D200000032000000010100000033000000060043006D000000000034000000010000000035000000045B903DA3000000440000000400000000000000450000000400000000000000460000000100000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	oent, err := chunk.NewOentChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewOentChunk err to be nil")
	}

	if oent.Header() != hdr {
		t.Fatal("expected header to be the same")
	}
}

func TestOentType(t *testing.T) {
	data, _ := hex.DecodeString("6F656E740000028F61646174000002870000000100000004000000D400000002000000CA002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D00700033000000000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D006900780029000000000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000000000090000000C0048006F00750073006500000000000F0000000400000077000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000000000170000000A003200300031003800000000001C000000045B903D080000001D000000045B903DA30000001F00000004000000010000002D000000040000009B0000003000000004000000D200000032000000010100000033000000060043006D000000000034000000010000000035000000045B903DA3000000440000000400000000000000450000000400000000000000460000000100000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	oent, err := chunk.NewOentChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewOentChunk err to be nil")
	}

	actual := oent.Type()
	expected := "oent"

	if actual != expected {
		t.Fatalf("expected type to be %s, got %s", expected, actual)
	}
}
