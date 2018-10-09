package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// FullPath ...
type FullPath struct {
	header *Header
	data   []byte
}

// Value ...
func (f *FullPath) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

// NewFullPathField ...
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
