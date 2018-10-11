package field

import (
	"encoding/binary"
	"io"
)

// PlayTime is the length of time the track was playing in Serato.
type PlayTime struct {
	header *Header
	data   []byte
}

// Value returns the play time.
// TODO: change to format MM:SS instead of int of seconds.
func (p *PlayTime) Value() int {
	return int(binary.BigEndian.Uint32(p.data))
}

// NewPlayTimeField returns a PlayTime, using the header to read the field data.
func NewPlayTimeField(header *Header, r io.Reader) (*PlayTime, error) {
	if header.Identifier != playtimeID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &PlayTime{header, data[:]}, nil
}
