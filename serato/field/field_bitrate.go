package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tombell/saga/strutil"
)

// Bitrate is the bitrate of the track.
type Bitrate struct {
	header *Header
	data   []byte
}

// Value returns the bitrate.
func (b *Bitrate) Value() string {
	s := strutil.DecodeUTF16(b.data)
	return strutil.TrimNull(s)
}

func (b *Bitrate) String() string {
	return fmt.Sprintf("Bitrate: %s", b.Value())
}

// NewBitrateField returns a Bitrate, using the header to read the field data.
func NewBitrateField(header *Header, r io.Reader) (*Bitrate, error) {
	if header.Identifier != bitrateID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Bitrate{header, data[:]}, nil
}
