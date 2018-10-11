package field

import (
	"encoding/binary"
	"io"
	"time"
)

// UpdatedAt is the date/time the track was updated at in Serato.
type UpdatedAt struct {
	header *Header
	data   []byte
}

// Value returns the updated at time.
func (u *UpdatedAt) Value() time.Time {
	ts := binary.BigEndian.Uint32(u.data)
	return time.Unix(int64(ts), 0).UTC()
}

// NewUpdatedAtField returns an UpdatedAt, using the header to read the field
// data.
func NewUpdatedAtField(header *Header, r io.Reader) (*UpdatedAt, error) {
	if header.Identifier != updatedAtID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &UpdatedAt{header, data[:]}, nil
}
