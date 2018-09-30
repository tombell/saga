package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const genreID = 9

// Genre ...
type Genre struct {
	header *Header
	data   []byte
}

// Value ...
func (g *Genre) Value() string {
	s := strutil.DecodeUTF16(g.data)
	return strutil.TrimNull(s)
}

// NewGenreField ...
func NewGenreField(header *Header, r io.Reader) (*Genre, error) {
	if header.Identifier != genreID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Genre{header, data[:]}, nil
}
