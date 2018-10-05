package chunk

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Header is a chunk header that contains information about the next chunk in
// the Serato session file format.
type Header struct {
	Identifier [4]byte
	Length     uint32
}

// Type returns the string representation of the chunk identifier.
func (h *Header) Type() string {
	return string(h.Identifier[:])
}

func (h *Header) String() string {
	return fmt.Sprintf("Chunk: %s, Data length: %d", h.Type(), h.Length)
}

// NewHeader returns a new Header that has been read from the given reader.
func NewHeader(r io.Reader) (*Header, error) {
	var hdr Header

	if err := binary.Read(r, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
