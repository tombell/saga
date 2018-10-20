package field

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// EndTime is the date/time the track finished playing in Serato.
type EndTime struct {
	header *Header
	data   []byte
}

// Value returns the end time.
func (f *EndTime) Value() time.Time {
	ts := binary.BigEndian.Uint32(f.data)
	return time.Unix(int64(ts), 0).UTC()
}

func (f *EndTime) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewEndTimeField returns an EndTime, using the header to read the field data.
func NewEndTimeField(header *Header, r io.Reader) (*EndTime, error) {
	if header.Identifier != endtimeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &EndTime{header, data[:]}, nil
}
