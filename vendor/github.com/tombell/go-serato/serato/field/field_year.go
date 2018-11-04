package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Year is the year of the track.
type Year struct {
	header *Header
	data   []byte
}

// Value returns the year.
func (f *Year) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Year) String() string {
	return f.Value()
}

// NewYearField returns an initialised Year, using the given field header.
func NewYearField(header *Header, r io.Reader) (*Year, error) {
	if header.Identifier != yearID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Year{
		header: header,
		data:   data[:],
	}, nil
}
