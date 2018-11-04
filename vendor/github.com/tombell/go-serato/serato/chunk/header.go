package chunk

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Header is a data structure that contains details about the following chunk in
// the Serato session file.
type Header struct {
	Identifier [4]byte
	Length     uint32
}

// Type returns a string of the identifier for the chunk.
func (h *Header) Type() string {
	return string(h.Identifier[:])
}

func (h *Header) String() string {
	return fmt.Sprintf("chunk (%s) length (%d)", h.Type(), h.Length)
}

// NewHeader returns an initialised Header that has been read using the given
// reader.
func NewHeader(r io.Reader) (*Header, error) {
	var h Header

	if err := binary.Read(r, binary.BigEndian, &h); err != nil {
		return nil, err
	}

	return &h, nil
}
