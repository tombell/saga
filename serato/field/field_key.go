package field

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/tombell/saga/strutil"
)

// TODO: Key is field #51

// Key ...
type Key struct {
	header *Header
	data   []byte
}

// Value ...
func (k *Key) Value() string {
	s := strutil.DecodeUTF16(k.data)
	return strings.Trim(s, string(0))
}

// NewKeyField ...
func NewKeyField(header *Header, r io.Reader) (*Key, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Key{header, data[:]}, nil
}
