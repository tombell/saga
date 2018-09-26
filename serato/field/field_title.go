package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Title is field #6

// Title ...
type Title struct {
	header *Header
	data   []byte
}

// Value ...
func (t *Title) Value() string {
	s := strutil.DecodeUTF16(t.data)
	return strings.Trim(s, string(0))
}

// NewTitleField ...
func NewTitleField(header *Header, r io.Reader) (*Title, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Title{header, data[:]}, nil
}
