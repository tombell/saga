package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// SessionID is the session ID of the track.
type SessionID struct {
	header *Header
	data   []byte
}

// Value returns the session ID.
func (s *SessionID) Value() int {
	return int(binary.BigEndian.Uint32(s.data))
}

func (s *SessionID) String() string {
	return fmt.Sprintf("Session ID: %d", s.Value())
}

// NewSessionIDField returns a SessionID, using the header to read the field
// data.
func NewSessionIDField(header *Header, r io.Reader) (*SessionID, error) {
	if header.Identifier != sessionID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &SessionID{header, data[:]}, nil
}
