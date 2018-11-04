package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// FullPath is the full path to the track on disk.
type FullPath struct {
	header *Header
	data   []byte
}

// Value returns the full path.
func (f *FullPath) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *FullPath) String() string {
	return f.Value()
}

// NewFullPathField returns an initialised FullPath, using the given field
// header.
func NewFullPathField(header *Header, r io.Reader) (*FullPath, error) {
	if header.Identifier != fullpathID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &FullPath{
		header: header,
		data:   data[:],
	}, nil
}
