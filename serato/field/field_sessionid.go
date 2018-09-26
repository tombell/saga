package field

import (
	"encoding/binary"
	"io"
)

// TODO: SessionID is field #48

// SessionID ...
type SessionID struct {
	header *Header
	data   []byte
}

// Value ...
func (s *SessionID) Value() int {
	return int(binary.BigEndian.Uint32(s.data))
}

// NewSessionIDField ...
func NewSessionIDField(header *Header, r io.Reader) (*SessionID, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &SessionID{header, data[:]}, nil
}
