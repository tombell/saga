package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Composer is the composer of the track.
type Composer struct {
	header *Header
	data   []byte
}

// Value returns the composer.
func (f *Composer) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Composer) String() string {
	return f.Value()
}

// NewComposerField returns a Composer, using the header to read the field data.
func NewComposerField(header *Header, r io.Reader) (*Composer, error) {
	if header.Identifier != composerID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Composer{header, data[:]}, nil
}
