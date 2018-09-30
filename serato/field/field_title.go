package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const titleID = 6

// Title ...
type Title struct {
	header *Header
	data   []byte
}

// Value ...
func (t *Title) Value() string {
	s := strutil.DecodeUTF16(t.data)
	return strutil.TrimNull(s)
}

// NewTitleField ...
func NewTitleField(header *Header, r io.Reader) (*Title, error) {
	if header.Identifier != titleID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Title{header, data[:]}, nil
}
