package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Language is the language of the track.
type Language struct {
	header *Header
	data   []byte
}

// Value returns the language.
func (f *Language) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Language) String() string {
	return f.Value()
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
