package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const bitrateID = 13

// Bitrate ...
type Bitrate struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Bitrate) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewBitrateField ...
func NewBitrateField(header *Header, r io.Reader) (*Artist, error) {
	if header.Identifier != bitrateID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Bitrate{header, data[:]}, nil
}
