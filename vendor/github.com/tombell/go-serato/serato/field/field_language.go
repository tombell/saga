package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/go-serato/internal/decode"
	"github.com/tombell/go-serato/internal/trim"
)

// Language is the language of the track.
type Language struct {
	header *Header
	data   []byte
}

// Value returns the language.
func (f *Language) Value() string {
	s := decode.UTF16(f.data)
	return trim.Nil(s)
}

func (f *Language) String() string {
	return f.Value()
}

// NewLanguageField returns an initialised Language, using the given field
// header.
func NewLanguageField(header *Header, r io.Reader) (*Language, error) {
	if header.Identifier != languageID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Language{
		header: header,
		data:   data[:],
	}, nil
}
