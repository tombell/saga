package field

import (
	"encoding/binary"
	"io"
)

const field72ID = 72

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
	if header.Identifier != field72ID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Field72{header, data[:]}, nil
}
