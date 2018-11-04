package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Bitrate is the bitrate of the track.
type Bitrate struct {
	header *Header
	data   []byte
}

// Value returns the bitrate.
func (f *Bitrate) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Bitrate) String() string {
	return f.Value()
}

// NewBitrateField returns an initialised Bitrate, using the given field header.
func NewBitrateField(header *Header, r io.Reader) (*Bitrate, error) {
	if header.Identifier != bitrateID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Bitrate{
		header: header,
		data:   data[:],
	}, nil
}
