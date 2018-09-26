package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Label is field #21

// Label ...
type Label struct {
	header *Header
	data   []byte
}

// Value ...
func (l *Label) Value() string {
	s := strutil.DecodeUTF16(l.data)
	return strings.Trim(s, string(0))
}

// NewLabelField ...
func NewLabelField(header *Header, r io.Reader) (*Label, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Label{header, data[:]}, nil
}
