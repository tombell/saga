package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// BPM is the BPM of the track.
type BPM struct {
	header *Header
	data   []byte
}

// Value returns the BPM.
func (f *BPM) Value() int {
	return int(binary.BigEndian.Uint32(f.data))
}

func (f *BPM) String() string {
	return fmt.Sprintf("%d", f.Value())
}

// NewBPMField returns an initialised BPM, using the given field header.
func NewBPMField(header *Header, r io.Reader) (*BPM, error) {
	if header.Identifier != bpmID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &BPM{
		header: header,
		data:   data[:],
	}, nil
}
