package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tombell/saga/strutil"
)

// Grouping is the grouping of the track.
type Grouping struct {
	header *Header
	data   []byte
}

// Value returns the grouping.
func (g *Grouping) Value() string {
	s := strutil.DecodeUTF16(g.data)
	return strutil.TrimNull(s)
}

func (g *Grouping) String() string {
	return fmt.Sprintf("Grouping: %s", g.Value())
}

// NewGroupingField returns a Grouping, using the header to read the field data.
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
