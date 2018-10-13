package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tombell/saga/strutil"
)

// Album is the album of the track.
type Album struct {
	header *Header
	data   []byte
}

// Value returns te album.
func (a *Album) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

func (a *Album) String() string {
	return fmt.Sprintf("Album: %s", a.Value())
}

// NewAlbumField returns an Album, using the header to read the field data.
func NewAlbumField(header *Header, r io.Reader) (*Album, error) {
	if header.Identifier != albumID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Album{header, data[:]}, nil
}
