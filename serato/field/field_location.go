package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Location is the base directory of the track.
type Location struct {
	header *Header
	data   []byte
}

// Value returns the location.
func (f *Location) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Location) String() string {
	return f.Value()
}

// NewLocationField returns a Location, using the header to read the field data.
func NewLocationField(header *Header, r io.Reader) (*Location, error) {
	if header.Identifier != locationID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Location{header, data[:]}, nil
}
