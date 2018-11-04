package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Size returns the size of the track file.
type Size struct {
	header *Header
	data   []byte
}

// Value returns the size.
func (f *Size) Value() string {
	str := decode.UTF16(f.data)
	return trim.Nil(str)
}

func (f *Size) String() string {
	return f.Value()
}

// NewSizeField returns an initialised Size, using the given field header.
func NewSizeField(header *Header, r io.Reader) (*Size, error) {
	if header.Identifier != sizeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Size{
		header: header,
		data:   data[:],
	}, nil
}
