package chunk

import (
	"encoding/binary"
	"io"
)

// Header ...
type Header struct {
	Identifier [4]byte
	Length     uint32
}

// Type ...
func (h *Header) Type() string {
	return string(h.Identifier[:])
}

// NewHeader ...
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
