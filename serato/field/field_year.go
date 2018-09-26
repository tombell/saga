package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Year is field #23

// Year ...
type Year struct {
	header *Header
	data   []byte
}

// Value ...
func (y *Year) Value() string {
	s := strutil.DecodeUTF16(y.data)
	return strings.Trim(s, string(0))
}

// NewYearField ...
func NewYearField(header *Header, r io.Reader) (*Year, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Year{header, data[:]}, nil
}
