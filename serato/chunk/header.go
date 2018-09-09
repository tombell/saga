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
func NewHeader(r io.Reader) *Header {
	var hdr Header
	binary.Read(r, binary.LittleEndian, &hdr)
	return &hdr
}
