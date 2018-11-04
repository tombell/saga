package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Field69 is the unknown field with ID 69.
type Field69 struct {
	header *Header
	data   []byte
}

// Value returns the raw byts for the field.
func (f *Field69) Value() []byte {
	return f.data
}

func (f *Field69) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewField69Field returns an initialised Field69, using the given field header.
func NewField69Field(header *Header, r io.Reader) (*Field69, error) {
	if header.Identifier != field69ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field69{
		header: header,
		data:   data[:],
	}, nil
}
