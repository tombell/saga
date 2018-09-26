package field

import (
	"encoding/binary"
	"io"
)

// TODO: Field72 is field #69

// Field72 ...
type Field72 struct {
	header *Header
	data   []byte
}

// Value ...
func (f *Field72) Value() []byte {
	return f.data
}

// NewField72Field ...
func NewField72Field(header *Header, r io.Reader) (*Field72, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field72{header, data[:]}, nil
}
