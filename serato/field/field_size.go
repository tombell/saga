package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tombell/saga/strutil"
)

// Size returns the size of the track file.
type Size struct {
	header *Header
	data   []byte
}

// Value returns the size.
func (s *Size) Value() string {
	str := strutil.DecodeUTF16(s.data)
	return strutil.TrimNull(str)
}

func (s *Size) String() string {
	return fmt.Sprintf("Size: %s", s.Value())
}

// NewSizeField returns a Size, using the header to read the field data.
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
