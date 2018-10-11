package field

import (
	"encoding/binary"
	"io"
)

// Row is the row of the track.
type Row struct {
	header *Header
	data   []byte
}

// Value returns the row.
func (r *Row) Value() int {
	return int(binary.BigEndian.Uint32(r.data))
}

// NewRowField returns a Row, using the header to read the field data.
func NewRowField(header *Header, r io.Reader) (*Row, error) {
	if header.Identifier != rowID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Row{header, data[:]}, nil
}
