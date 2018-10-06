package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const lengthID = 10

// Length ...
type Length struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Length) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewLengthField ...
func NewLengthField(header *Header, r io.Reader) (*Artist, error) {
	if header.Identifier != lengthID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Length{header, data[:]}, nil
}
