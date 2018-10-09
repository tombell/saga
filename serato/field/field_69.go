package field

import (
	"encoding/binary"
	"io"
)

// Field69 ...
type Field69 struct {
	header *Header
	data   []byte
}

// Value ...
func (f *Field69) Value() []byte {
	return f.data
}

// NewField69Field ...
func NewField69Field(header *Header, r io.Reader) (*Field69, error) {
	if header.Identifier != field69ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field69{header, data[:]}, nil
}
