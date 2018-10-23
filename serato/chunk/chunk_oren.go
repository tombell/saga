package chunk

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Oren is a chunk that contains a single UENT chunk which contains a reference
// to an OENT chunk to be deleted.
type Oren struct {
	header *Header
	data   []byte

	Uent *Uent
}

// Header returns the header of the chunk.
func (o *Oren) Header() *Header {
	return o.header
}

// Type returns the type of the chunk.
func (o *Oren) Type() string {
	return o.header.Type()
}

// NewOrenChunk returns an OREN chunk, using the header to read the chunk data.
func NewOrenChunk(header *Header, r io.Reader) (*Oren, error) {
	if header.Type() != orenID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(data[:])
	hdr, _ := NewHeader(buf)

	uent, err := NewUentChunk(hdr, buf)
	if err != nil {
		return nil, err
	}

	return &Oren{
		header: header,
		data:   data[:],
		Uent:   uent,
	}, nil
}
