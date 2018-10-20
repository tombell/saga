package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Key is the key of the track.
type Key struct {
	header *Header
	data   []byte
}

// Value returns the key.
func (f *Key) Value() string {
	s := strutil.DecodeUTF16(f.data)
	return strutil.TrimNull(s)
}

func (f *Key) String() string {
	return f.Value()
}

// NewKeyField returns a Key, using the header to read the field data.
func NewKeyField(header *Header, r io.Reader) (*Key, error) {
	if header.Identifier != keyID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Key{header, data[:]}, nil
}
