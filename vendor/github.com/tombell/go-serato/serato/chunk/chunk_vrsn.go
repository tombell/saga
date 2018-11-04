package chunk

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
)

// Vrsn is a chunk that contains the version of the Serato session file format.
type Vrsn struct {
	header *Header
	data   []byte
}

// Header returns the header for the chunk.
func (v *Vrsn) Header() *Header {
	return v.header
}

// Type returns the type of the chunk.
func (v *Vrsn) Type() string {
	return v.header.Type()
}

// Version returns the version string of the Serato session file format.
func (v *Vrsn) Version() string {
	return decode.UTF16(v.data)
}

// NewVrsnChunk returns an initialised Vrsn chunk, using the header to read the
// chunk data.
func NewVrsnChunk(header *Header, r io.Reader) (*Vrsn, error) {
	if header.Type() != vrsnID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Vrsn{
		header: header,
		data:   data[:],
	}, nil
}
