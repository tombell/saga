package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Row is the row of the track.
type Row struct {
	header *Header
	data   []byte
}

// Value returns the row.
func (f *Row) Value() int {
	return int(binary.BigEndian.Uint32(f.data))
}

func (f *Row) String() string {
	return fmt.Sprintf("%d", f.Value())
}

// NewRowField returns an initialised Row using the given field header.
func NewRowField(header *Header, r io.Reader) (*Row, error) {
	if header.Identifier != rowID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Row{
		header: header,
		data:   data[:],
	}, nil
}
