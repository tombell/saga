package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Title is the title of the track.
type Title struct {
	header *Header
	data   []byte
}

// Value returns the title.
func (f *Title) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Title) String() string {
	return f.Value()
}

// NewTitleField returns an initialised Title, using the given field header.
func NewTitleField(header *Header, r io.Reader) (*Title, error) {
	if header.Identifier != titleID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Title{
		header: header,
		data:   data[:],
	}, nil
}
