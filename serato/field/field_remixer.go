package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Remixer is the remixer of the track.
type Remixer struct {
	header *Header
	data   []byte
}

// Value returns the remixer.
func (f *Remixer) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Remixer) String() string {
	return f.Value()
}

// NewRemixerField returns a Remixer, using the header to read the field data.
func NewRemixerField(header *Header, r io.Reader) (*Remixer, error) {
	if header.Identifier != remixerID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Remixer{header, data[:]}, nil
}
