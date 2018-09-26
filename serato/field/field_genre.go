package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Genre is field #9

// Genre ...
type Genre struct {
	header *Header
	data   []byte
}

// Value ...
func (g *Genre) Value() string {
	s := strutil.DecodeUTF16(g.data)
	return strings.Trim(s, string(0))
}

// NewGenreField ...
func NewGenreField(header *Header, r io.Reader) (*Genre, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Genre{header, data[:]}, nil
}
