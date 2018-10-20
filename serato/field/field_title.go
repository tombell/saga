package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Title is the title of the track.
type Title struct {
	header *Header
	data   []byte
}

// Value returns the title.
func (f *Title) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Title) String() string {
	return f.Value()
}

// NewTitleField returns a Title, using the header to read the field data.
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
