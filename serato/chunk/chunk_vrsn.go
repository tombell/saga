package chunk

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Vrsn represents a vrsn chunk from a Serato session file.
type Vrsn struct {
	header *Header
	data   []byte
}

// Header returns the header of the chunk.
func (v *Vrsn) Header() *Header {
	return v.header
}

// Type returns the type of the chunk.
func (v *Vrsn) Type() string {
	return v.header.Type()
}

// Version returns the version of the Serato session file format.
func (v *Vrsn) Version() string {
	return strutil.DecodeUTF16(v.data)
}

// NewVrsnChunk returns a new vrsn chunk, using the header data to read the vrsn
// chunk data.
func NewVrsnChunk(header *Header, r io.Reader) (*Vrsn, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Vrsn{header, data[:]}, nil
}
