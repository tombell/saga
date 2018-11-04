package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Field72 is the unknown field with ID 72.
type Field72 struct {
	header *Header
	data   []byte
}

// Value returns the raw bytes for the field.
func (f *Field72) Value() []byte {
	return f.data
}

func (f *Field72) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewField72Field returns an initialised Field72, using the given field header.
func NewField72Field(header *Header, r io.Reader) (*Field72, error) {
	if header.Identifier != field72ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field72{
		header: header,
		data:   data[:],
	}, nil
}
