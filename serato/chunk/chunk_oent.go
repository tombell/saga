package chunk

import (
	"encoding/binary"
	"io"
)

// Oent is a chunk that contains a single ADAT chunk which contains track
// information.
type Oent struct {
	header *Header
	data   []byte

	// Adat is a single ADAT chunk, containing the track information.
	Adat *Adat
}

// Header returns the header of the chunk.
func (o *Oent) Header() *Header {
	return o.header
}

// Type returns the type of the chunk.
func (o *Oent) Type() string {
	return o.header.Type()
}

// NewOentChunk returns an OENT chunk, using the header to read the chunk data.
func NewOentChunk(header *Header, r io.Reader) (*Oent, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Oent{header, data[:], nil}, nil
}
