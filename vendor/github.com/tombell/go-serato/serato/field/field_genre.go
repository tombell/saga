package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Genre is the genre of the track.
type Genre struct {
	header *Header
	data   []byte
}

// Value returns the genre.
func (f *Genre) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Genre) String() string {
	return f.Value()
}

// NewGenreField returns an initialised Genre, using the given field header.
func NewGenreField(header *Header, r io.Reader) (*Genre, error) {
	if header.Identifier != genreID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Genre{
		header: header,
		data:   data[:],
	}, nil
}
