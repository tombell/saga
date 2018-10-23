package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Added is the added status of the track in Serato.
type Added struct {
	header *Header
	data   []byte
}

// Value returns the added status.
func (f *Added) Value() bool {
	return f.data[0] == byte(1)
}

func (f *Added) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewAddedField returns an Added, using the header to read the field data.
func NewAddedField(header *Header, r io.Reader) (*Added, error) {
	if header.Identifier != addedID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Added{header, data[:]}, nil
}
