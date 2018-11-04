package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Frequency is the frequency of the track.
type Frequency struct {
	header *Header
	data   []byte
}

// Value returns the frequency.
func (f *Frequency) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Frequency) String() string {
	return f.Value()
}

// NewFrequencyField returns an initialised Frequency, using the given field
// header.
func NewFrequencyField(header *Header, r io.Reader) (*Frequency, error) {
	if header.Identifier != frequencyID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Frequency{
		header: header,
		data:   data[:],
	}, nil
}
