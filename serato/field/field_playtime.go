package field

import (
	"encoding/binary"
	"io"
)

const playtimeID = 45

// PlayTime ...
type PlayTime struct {
	header *Header
	data   []byte
}

// Value ...
// TODO: change to format MM:SS instead of int of seconds.
func (p *PlayTime) Value() int {
	return int(binary.BigEndian.Uint32(p.data))
}

// NewPlayTimeField ...
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
