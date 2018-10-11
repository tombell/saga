package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Label is the label of the track.
type Label struct {
	header *Header
	data   []byte
}

// Value returns the label.
func (l *Label) Value() string {
	s := strutil.DecodeUTF16(l.data)
	return strutil.TrimNull(s)
}

// NewLabelField returns a Label, using the header to read the field data.
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
