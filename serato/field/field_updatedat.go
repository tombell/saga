package field

import (
	"encoding/binary"
	"io"
	"time"
)

// TODO: UpdatedAt is field #53

// UpdatedAt ...
type UpdatedAt struct {
	header *Header
	data   []byte
}

// Value ...
func (u *UpdatedAt) Value() time.Time {
	ts := binary.BigEndian.Uint32(u.data)
	return time.Unix(int64(ts), 0).UTC()
}

// NewUpdatedAtField ...
func NewUpdatedAtField(header *Header, r io.Reader) (*UpdatedAt, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &UpdatedAt{header, data[:]}, nil
}
