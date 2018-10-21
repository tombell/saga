package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Grouping is the grouping of the track.
type Grouping struct {
	header *Header
	data   []byte
}

// Value returns the grouping.
func (f *Grouping) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Grouping) String() string {
	return f.Value()
}

// NewGroupingField returns a Grouping, using the header to read the field data.
func NewGroupingField(header *Header, r io.Reader) (*Grouping, error) {
	if header.Identifier != groupingID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Grouping{header, data[:]}, nil
}
