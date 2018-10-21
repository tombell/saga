package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Label is the label of the track.
type Label struct {
	header *Header
	data   []byte
}

// Value returns the label.
func (f *Label) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Label) String() string {
	return f.Value()
}

// NewLabelField returns a Label, using the header to read the field data.
func NewLabelField(header *Header, r io.Reader) (*Label, error) {
	if header.Identifier != labelID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Label{header, data[:]}, nil
}
