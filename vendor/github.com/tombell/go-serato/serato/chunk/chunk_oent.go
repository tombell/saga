package chunk

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Oent is a chunk that contains a single ADAT chunk which contains track
// information.
type Oent struct {
	header *Header
	data   []byte

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
	if header.Type() != oentID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(data[:])
	hdr, _ := NewHeader(buf)

	adat, err := NewAdatChunk(hdr, buf)
	if err != nil {
		return nil, err
	}

	return &Oent{
		header: header,
		data:   data[:],
		Adat:   adat,
	}, nil
}
