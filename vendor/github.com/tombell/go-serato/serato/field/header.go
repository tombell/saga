package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Header is a data structure that contains details about the following field in
// the ADAT chunk, in the Serato session file.
type Header struct {
	Identifier uint32
	Length     uint32
}

func (h *Header) String() string {
	return fmt.Sprintf("field (%d) length (%d)", h.Identifier, h.Length)
}

// NewHeader returns an initialised Header that has been read using the given
// reader.
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
