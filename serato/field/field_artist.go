package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Artist is the artist of the track.
type Artist struct {
	header *Header
	data   []byte
}

// Value returns the artist.
func (f *Artist) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Artist) String() string {
	return f.Value()
}

// NewArtistField returns a Title, using the header to read the field data.
func NewArtistField(header *Header, r io.Reader) (*Artist, error) {
	if header.Identifier != artistID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Artist{header, data[:]}, nil
}
