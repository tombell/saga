package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Length is the length of the track.
type Length struct {
	header *Header
	data   []byte
}

// Value returns the length.
func (f *Length) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Length) String() string {
	return f.Value()
}

// NewLengthField returns an initialised Length, using the given field header.
func NewLengthField(header *Header, r io.Reader) (*Length, error) {
	if header.Identifier != lengthID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Length{
		header: header,
		data:   data[:],
	}, nil
}
