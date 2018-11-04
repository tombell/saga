package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Filename is the file name of the track.
type Filename struct {
	header *Header
	data   []byte
}

// Value returns the file name.
func (f *Filename) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Filename) String() string {
	return f.Value()
}

// NewFilenameField returns an initialised Filename, using the given field
// header.
func NewFilenameField(header *Header, r io.Reader) (*Filename, error) {
	if header.Identifier != filenameID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Filename{
		header: header,
		data:   data[:],
	}, nil
}
