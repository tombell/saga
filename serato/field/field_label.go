package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Label ...
type Label struct {
	header *Header
	data   []byte
}

// Value ...
func (l *Label) Value() string {
	s := strutil.DecodeUTF16(l.data)
	return strutil.TrimNull(s)
}

// NewLabelField ...
func NewLabelField(header *Header, r io.Reader) (*Label, error) {
	if header.Identifier != labelID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Label{header, data[:]}, nil
}
