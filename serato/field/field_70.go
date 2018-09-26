package field

import (
	"encoding/binary"
	"io"
)

// TODO: Field70 is field #70

// Field70 ...
type Field70 struct {
	header *Header
	data   []byte
}

// Value ...
func (f *Field70) Value() byte {
	return f.data[0]
}

// NewField70Field ...
func NewField70Field(header *Header, r io.Reader) (*Field70, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field70{header, data[:]}, nil
}
