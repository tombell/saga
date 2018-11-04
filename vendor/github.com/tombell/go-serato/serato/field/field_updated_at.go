package field

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// UpdatedAt is the date/time the track was updated at in Serato.
type UpdatedAt struct {
	header *Header
	data   []byte
}

// Value returns the updated at time.
func (f *UpdatedAt) Value() time.Time {
	ts := binary.BigEndian.Uint32(f.data)
	return time.Unix(int64(ts), 0).UTC()
}

func (f *UpdatedAt) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewUpdatedAtField returns an initialised UpdatedAt, using the given field
// header.
func NewUpdatedAtField(header *Header, r io.Reader) (*UpdatedAt, error) {
	if header.Identifier != updatedAtID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &UpdatedAt{
		header: header,
		data:   data[:],
	}, nil
}
