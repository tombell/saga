package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Year is the year of the track.
type Year struct {
	header *Header
	data   []byte
}

// Value returns the year.
func (f *Year) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Year) String() string {
	return f.Value()
}

// NewYearField returns a Year, using the header to read the field data.
func NewYearField(header *Header, r io.Reader) (*Year, error) {
	if header.Identifier != yearID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Year{header, data[:]}, nil
}
