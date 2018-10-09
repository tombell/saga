package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Size ...
type Size struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Size) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewSizeField ...
func NewSizeField(header *Header, r io.Reader) (*Size, error) {
	if header.Identifier != sizeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Size{header, data[:]}, nil
}
