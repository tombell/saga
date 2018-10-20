package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Length is the length of the track.
type Length struct {
	header *Header
	data   []byte
}

// Value returns the length.
func (f *Length) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Length) String() string {
	return f.Value()
}

// NewLengthField returns a Length, using the header to read the field data.
func NewLengthField(header *Header, r io.Reader) (*Length, error) {
	if header.Identifier != lengthID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Length{header, data[:]}, nil
}
