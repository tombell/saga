package chunk

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/serato/field"
)

// Adat is a chunk that contains different fields for track information.
type Adat struct {
	header *Header
	data   []byte

	*field.Fields
}

// Header returns the header of the chunk.
func (a *Adat) Header() *Header {
	return a.header
}

// Type returns the type of the chunk.
func (a *Adat) Type() string {
	return a.header.Type()
}

// NewAdatChunk returns an ADAT chunk, using the header to read the chunk data.
func NewAdatChunk(header *Header, r io.Reader) (*Adat, error) {
	if header.Type() != adatID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	fields, err := field.NewFields(data)
	if err != nil {
		return nil, err
	}

	return &Adat{
		header: header,
		data:   data[:],
		Fields: fields,
	}, nil
}
