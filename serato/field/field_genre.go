package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Genre is the genre of the track.
type Genre struct {
	header *Header
	data   []byte
}

// Value returns the genre.
func (f *Genre) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Genre) String() string {
	return f.Value()
}

// NewGenreField returns a Genre, using the header to read the field data.
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
