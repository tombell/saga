package field

import (
	"encoding/binary"
	"io"
)

// Row ...
type Row struct {
	header *Header
	data   []byte
}

// Value ...
func (r *Row) Value() int {
	return int(binary.BigEndian.Uint32(r.data))
}

// NewRowField ...
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
