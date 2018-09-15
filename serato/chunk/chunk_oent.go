package chunk

import (
	"encoding/binary"
	"io"
)

// TODO: Session files have these, each containing a single ADAT for a track

// Oent represents a oent chunk from a Serato session file.
type Oent struct {
	header *Header
	data   []byte
	// TODO: Single ADAT chunk for a track
}

// Header returns the header of the chunk.
func (o *Oent) Header() *Header {
	return o.header
}

// Type returns the type of the chunk.
func (o *Oent) Type() string {
	return o.header.Type()
}

// NewOentChunk returns a new oent chunk, using the header data to read the oent
// chunk data.
func NewOentChunk(header *Header, r io.Reader) (*Oent, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Oent{header, data[:]}, nil
}
