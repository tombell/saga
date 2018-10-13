package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Field39 is the unknown field with ID 39.
type Field39 struct {
	header *Header
	data   []byte
}

// Value returns the raw bytes for the field.
func (f *Field39) Value() []byte {
	return f.data
}

func (f *Field39) String() string {
	return fmt.Sprintf("Field 39: %v", f.Value())
}

// NewField39Field returns a Field39, using the header to read the field data.
func NewField39Field(header *Header, r io.Reader) (*Field39, error) {
	if header.Identifier != field39ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field39{header, data[:]}, nil
}
