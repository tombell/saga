package chunk

import (
	"encoding/binary"
	"io"
)

// Header represents the header for a chunk in the Serato session file format.
type Header struct {
	Identifier [4]byte
	Length     uint32
}

// Type returns the string representation of the chunk type.
func (h *Header) Type() string {
	return string(h.Identifier[:])
}

// TODO: implement for debugging purposes.
func (h *Header) String() string {
	return ""
}

// NewHeader returns an initialised Header by reading the next header.
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
