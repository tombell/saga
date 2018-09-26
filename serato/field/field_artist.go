package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Artist is field #7

// Artist ...
type Artist struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Artist) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strings.Trim(s, string(0))
}

// NewArtistField ...
func NewArtistField(header *Header, r io.Reader) (*Artist, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Artist{header, data[:]}, nil
}
