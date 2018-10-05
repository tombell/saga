package chunk

import (
	"encoding/binary"
	"fmt"
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

func (h *Header) String() string {
	return fmt.Sprintf("Chunk: %s, Data length: %d", h.Type(), h.Length)
}

// NewHeader returns an initialised Header by reading the next header.
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
