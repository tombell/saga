package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const yearID = 23

// Year ...
type Year struct {
	header *Header
	data   []byte
}

// Value ...
func (y *Year) Value() string {
	s := strutil.DecodeUTF16(y.data)
	return strutil.TrimNull(s)
}

// NewYearField ...
func NewYearField(header *Header, r io.Reader) (*Year, error) {
	if header.Identifier != yearID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Year{header, data[:]}, nil
}
