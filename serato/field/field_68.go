package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Field68 is the unknown field with ID 68.
type Field68 struct {
	header *Header
	data   []byte
}

// Value returns the raw bytes for the field.
func (f *Field68) Value() []byte {
	return f.data
}

func (f *Field68) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewField68Field returns a Field68, using the header to read the field data.
func NewField68Field(header *Header, r io.Reader) (*Field68, error) {
	if header.Identifier != field68ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field68{header, data[:]}, nil
}
