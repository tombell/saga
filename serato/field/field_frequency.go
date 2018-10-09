package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Frequency ...
type Frequency struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Frequency) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewFrequencyField ...
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
