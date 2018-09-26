package field

import (
	"encoding/binary"
	"io"
)

// TODO: Added is field #52

// Added ...
type Added struct {
	header *Header
	data   []byte
}

// Value ...
func (p *Added) Value() byte {
	return p.data[0]
}

// NewAddedField ...
func NewAddedField(header *Header, r io.Reader) (*Added, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Added{header, data[:]}, nil
}
