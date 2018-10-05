package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Header represents the header for a field in an adat chunk in a Serato session
// file format.
type Header struct {
	Identifier uint32
	Length     uint32
}

func (h *Header) String() string {
	return fmt.Sprintf("Field: %d, Data length: %d", h.Identifier, h.Length)
}

// NewHeader returns an initialised Header by reading the next header.
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
