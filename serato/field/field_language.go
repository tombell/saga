package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Language is the language of the track.
type Language struct {
	header *Header
	data   []byte
}

// Value returns the language.
func (l *Language) Value() string {
	s := strutil.DecodeUTF16(l.data)
	return strutil.TrimNull(s)
}

func (l *Language) String() string {
	return l.Value()
}

// NewLanguageField returns a Language, using the header to read the field data.
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
