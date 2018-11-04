package chunk

import (
	"encoding/binary"
	"io"
)

// Uent is a chunk that contains an identifier for a previous OENT chunk to be
// deleted.
type Uent struct {
	header *Header
	data   []byte
}

// Header returns the header of the chunk.
func (u *Uent) Header() *Header {
	return u.header
}

// Type returns the type of the chunk.
func (u *Uent) Type() string {
	return u.header.Type()
}

// Value returns the row of the OENT to be deleted.
func (u *Uent) Value() int {
	return int(binary.BigEndian.Uint32(u.data))
}

// NewUentChunk returns an UENT chunk, using the header to read the chunk data.
func NewUentChunk(header *Header, r io.Reader) (*Uent, error) {
	if header.Type() != uentID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Uent{
		header: header,
		data:   data[:],
	}, nil
}
