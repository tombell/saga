package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tombell/saga/strutil"
)

// Frequency is the frequency of the track.
type Frequency struct {
	header *Header
	data   []byte
}

// Value returns the frequency.
func (f *Frequency) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Frequency) String() string {
	return fmt.Sprintf("Frequency: %s", f.Value())
}

// NewFrequencyField returns a Frequency, using the header to read the field
// data.
func NewFrequencyField(header *Header, r io.Reader) (*Frequency, error) {
	if header.Identifier != frequencyID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Frequency{header, data[:]}, nil
}
