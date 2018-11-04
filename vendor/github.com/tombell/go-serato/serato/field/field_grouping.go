package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Grouping is the grouping of the track.
type Grouping struct {
	header *Header
	data   []byte
}

// Value returns the grouping.
func (f *Grouping) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Grouping) String() string {
	return f.Value()
}

// NewGroupingField returns an initialised Grouping, using the given field
// header.
func NewGroupingField(header *Header, r io.Reader) (*Grouping, error) {
	if header.Identifier != groupingID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Grouping{
		header: header,
		data:   data[:],
	}, nil
}
