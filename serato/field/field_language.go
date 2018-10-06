package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const languageID = 18

// Language ...
type Language struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Language) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewLanguageField ...
func NewLanguageField(header *Header, r io.Reader) (*Language, error) {
	if header.Identifier != languageID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Language{header, data[:]}, nil
}
