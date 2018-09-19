package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: FullPath is field #2

// FullPath ...
type FullPath struct {
	header *Header
	data   []byte
}

// Value ...
func (f *FullPath) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strings.Trim(s, string(0))
}

// NewFullPathField ...
func NewFullPathField(header *Header, r io.Reader) (*FullPath, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &FullPath{header, data[:]}, nil
}
