package field

import (
	"encoding/binary"
	"io"
	"time"
)

const starttimeID = 28

// StartTime ...
type StartTime struct {
	header *Header
	data   []byte
}

// Value ...
func (s *StartTime) Value() time.Time {
	ts := binary.BigEndian.Uint32(s.data)
	return time.Unix(int64(ts), 0).UTC()
}

// NewStartTimeField ...
func NewStartTimeField(header *Header, r io.Reader) (*StartTime, error) {
	if header.Identifier != starttimeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &StartTime{header, data[:]}, nil
}
