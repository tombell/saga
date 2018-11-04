package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Album is the album of the track.
type Album struct {
	header *Header
	data   []byte
}

// Value returns te album.
func (f *Album) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Album) String() string {
	return f.Value()
}

// NewAlbumField returns an initialised Album, using the given field header.
func NewAlbumField(header *Header, r io.Reader) (*Album, error) {
	if header.Identifier != albumID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Album{
		header: header,
		data:   data[:],
	}, nil
}
