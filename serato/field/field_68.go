package field

import (
	"encoding/binary"
	"io"
)

// TODO: Field68 is field #68

// Field68 ...
type Field68 struct {
	header *Header
	data   []byte
}

// Value ...
func (f *Field68) Value() []byte {
	return f.data
}

// NewField68Field ...
func NewField68Field(header *Header, r io.Reader) (*Field68, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field68{header, data[:]}, nil
}
