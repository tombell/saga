package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tombell/saga/strutil"
)

// Location is the base directory of the track.
type Location struct {
	header *Header
	data   []byte
}

// Value returns the location.
func (l *Location) Value() string {
	s := strutil.DecodeUTF16(l.data)
	return strutil.TrimNull(s)
}

func (l *Location) String() string {
	return fmt.Sprintf("Location: %s", l.Value())
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
