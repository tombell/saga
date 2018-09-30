package field

import (
	"encoding/binary"
	"io"
	"time"
)

const endtimeID = 29

// EndTime ...
type EndTime struct {
	header *Header
	data   []byte
}

// Value ...
func (e *EndTime) Value() time.Time {
	ts := binary.BigEndian.Uint32(e.data)
	return time.Unix(int64(ts), 0).UTC()
}

// NewEndTimeField ...
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
