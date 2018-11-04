package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Remixer is the remixer of the track.
type Remixer struct {
	header *Header
	data   []byte
}

// Value returns the remixer.
func (f *Remixer) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Remixer) String() string {
	return f.Value()
}

// NewRemixerField returns an initialised Remixer, using the given field header.
func NewRemixerField(header *Header, r io.Reader) (*Remixer, error) {
	if header.Identifier != remixerID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Remixer{
		header: header,
		data:   data[:],
	}, nil
}
