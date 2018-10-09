package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Composer ...
type Composer struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Composer) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewComposerField ...
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
