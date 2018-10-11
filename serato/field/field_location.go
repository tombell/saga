package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Location is the base directory of the track.
type Location struct {
	header *Header
	data   []byte
}

// Value returns the location.
func (a *Location) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
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
