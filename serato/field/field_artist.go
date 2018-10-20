package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Artist is the artist of the track.
type Artist struct {
	header *Header
	data   []byte
}

// Value returns the artist.
func (f *Artist) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
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
