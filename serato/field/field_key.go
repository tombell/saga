package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

const keyID = 51

// Key ...
type Key struct {
	header *Header
	data   []byte
}

// Value ...
func (k *Key) Value() string {
	s := strutil.DecodeUTF16(k.data)
	return strutil.TrimNull(s)
}

// NewKeyField ...
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
