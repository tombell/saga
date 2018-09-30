package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const groupingID = 19

// Grouping ...
type Grouping struct {
	header *Header
	data   []byte
}

// Value ...
func (g *Grouping) Value() string {
	s := strutil.DecodeUTF16(g.data)
	return strutil.TrimNull(s)
}

// NewGroupingField ...
func NewGroupingField(header *Header, r io.Reader) (*Grouping, error) {
	if header.Identifier != groupingID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Grouping{header, data[:]}, nil
}
