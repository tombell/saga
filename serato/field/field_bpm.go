package field

import (
	"encoding/binary"
	"io"
)

// BPM is the BPM of the track.
type BPM struct {
	header *Header
	data   []byte
}

// Value returns the BPM.
func (b *BPM) Value() int {
	return int(binary.BigEndian.Uint32(b.data))
}

func (b *BPM) String() string {
	return b.Value()
}

// NewBPMField returns a BPM, using the header to read the field data.
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
