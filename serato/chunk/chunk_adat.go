package chunk

import (
	"encoding/binary"
	"io"
)

// TODO: Data chunk, contains fields. Fields meaning file format dependent

// Adat represents an adat chunk from a Serato session file.
type Adat struct {
	header *Header
	data   []byte
	// TODO: Fields for the ADAT data
}

// Header returns the header of the chunk.
func (a *Adat) Header() *Header {
	return a.header
}

// Type returns the type of the chunk.
func (a *Adat) Type() string {
	return a.header.Type()
}

// NewAdatChunk returns a new adat chunk, using the header data to read the adat
// chunk data.
func NewAdatChunk(header *Header, r io.Reader) (*Adat, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Adat{header, data[:]}, nil
}
