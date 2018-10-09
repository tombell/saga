package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Remixer ...
type Remixer struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Remixer) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewRemixerField ...
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
