package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Grouping is field #19

// Grouping ...
type Grouping struct {
	header *Header
	data   []byte
}

// Value ...
func (g *Grouping) Value() string {
	s := strutil.DecodeUTF16(g.data)
	return strings.Trim(s, string(0))
}

// NewGroupingField ...
func NewGroupingField(header *Header, r io.Reader) (*Grouping, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Grouping{header, data[:]}, nil
}
