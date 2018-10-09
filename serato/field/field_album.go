package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Album ...
type Album struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Album) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewAlbumField ...
func NewAlbumField(header *Header, r io.Reader) (*Artist, error) {
	if header.Identifier != albumID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Artist{header, data[:]}, nil
}
