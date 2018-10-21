package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Size returns the size of the track file.
type Size struct {
	header *Header
	data   []byte
}

// Value returns the size.
func (f *Size) Value() string {
	str := decode.UTF16(f.data)
	return trim.Null(str)
}

func (f *Size) String() string {
	return f.Value()
}

// NewSizeField returns a Size, using the header to read the field data.
func NewSizeField(header *Header, r io.Reader) (*Size, error) {
	if header.Identifier != sizeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Size{header, data[:]}, nil
}
