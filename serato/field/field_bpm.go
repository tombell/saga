package field

import (
	"encoding/binary"
	"io"
)

const bpmID = 15

// BPM ...
type BPM struct {
	header *Header
	data   []byte
}

// Value ...
func (b *BPM) Value() int {
	return int(binary.BigEndian.Uint32(b.data))
}

// NewBPMField ...
func NewBPMField(header *Header, r io.Reader) (*BPM, error) {
	if header.Identifier != bpmID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &BPM{header, data[:]}, nil
}
