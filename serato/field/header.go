package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Header is a field header that contains information about the next field in
// the ADAT chunk, as part of the Serato file format.
type Header struct {
	Identifier uint32
	Length     uint32
}

func (h *Header) String() string {
	return fmt.Sprintf("Field: %d, Data length: %d", h.Identifier, h.Length)
}

// NewHeader returns a new Header that has been read from the given reader.
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
