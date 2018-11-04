package field

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// StartTime is the date/time the track started playing in Serato.
type StartTime struct {
	header *Header
	data   []byte
}

// Value returns the start time.
func (f *StartTime) Value() time.Time {
	ts := binary.BigEndian.Uint32(f.data)
	return time.Unix(int64(ts), 0).UTC()
}

func (f *StartTime) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewStartTimeField returns an initialised StartTime, using the given field
// header.
func NewStartTimeField(header *Header, r io.Reader) (*StartTime, error) {
	if header.Identifier != starttimeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &StartTime{
		header: header,
		data:   data[:],
	}, nil
}
