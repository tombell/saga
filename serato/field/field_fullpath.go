package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// FullPath is the full path of the track.
type FullPath struct {
	header *Header
	data   []byte
}

// Value returns the full path.
func (f *FullPath) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

// NewFullPathField returns a FullPath, using the header to read the field data.
func NewFullPathField(header *Header, r io.Reader) (*FullPath, error) {
	if header.Identifier != fullpathID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &FullPath{header, data[:]}, nil
}
