package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// FullPath is the full path of the track.
type FullPath struct {
	header *Header
	data   []byte
}

// Value returns the full path.
func (f *FullPath) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *FullPath) String() string {
	return f.Value()
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
