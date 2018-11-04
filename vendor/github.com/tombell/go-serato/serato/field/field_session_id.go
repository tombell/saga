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
func (f *SessionID) Value() int {
	return int(binary.BigEndian.Uint32(f.data))
}

func (f *SessionID) String() string {
	return fmt.Sprintf("%d", f.Value())
}

// NewSessionIDField returns an initialised SessionID, using the given field
// header.
func NewSessionIDField(header *Header, r io.Reader) (*SessionID, error) {
	if header.Identifier != sessionID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &SessionID{
		header: header,
		data:   data[:],
	}, nil
}
