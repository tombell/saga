package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const locationID = 3

// Location ...
type Location struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Location) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewLocationField ...
func NewLocationField(header *Header, r io.Reader) (*Artist, error) {
	if header.Identifier != locationID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Location)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Location{header, data[:]}, nil
}
