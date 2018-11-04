package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PlayTime is the length of time the track was playing in Serato.
type PlayTime struct {
	header *Header
	data   []byte
}

// Value returns the play time.
func (f *PlayTime) Value() int {
	return int(binary.BigEndian.Uint32(f.data))
}

func (f *PlayTime) String() string {
	return fmt.Sprintf("%d", f.Value())
}

// NewPlayTimeField returns an initialised PlayTime, using the given field
// header.
func NewPlayTimeField(header *Header, r io.Reader) (*PlayTime, error) {
	if header.Identifier != playtimeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &PlayTime{
		header: header,
		data:   data[:],
	}, nil
}
