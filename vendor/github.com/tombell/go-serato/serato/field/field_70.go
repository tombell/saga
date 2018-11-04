package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Field70 is the unknown field with ID 70.
type Field70 struct {
	header *Header
	data   []byte
}

// Value returns the raw bytes for the field.
func (f *Field70) Value() byte {
	return f.data[0]
}

func (f *Field70) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewField70Field returns an initialised Field70, using the given field header.
func NewField70Field(header *Header, r io.Reader) (*Field70, error) {
	if header.Identifier != field70ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field70{
		header: header,
		data:   data[:],
	}, nil
}
