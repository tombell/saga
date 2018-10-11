package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Filename is the file name of the track.
type Filename struct {
	header *Header
	data   []byte
}

// Value returns the file name.
func (a *Filename) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewFilenameField returns a Filename, using the header to read the field data.
func NewFilenameField(header *Header, r io.Reader) (*Filename, error) {
	if header.Identifier != filenameID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Filename{header, data[:]}, nil
}
