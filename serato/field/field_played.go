package field

import (
	"encoding/binary"
	"io"
)

const playedID = 50

// Played ...
type Played struct {
	header *Header
	data   []byte
}

// Value ...
func (p *Played) Value() byte {
	return p.data[0]
}

// NewPlayedField ...
func NewPlayedField(header *Header, r io.Reader) (*Played, error) {
	if header.Identifier != playedID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Played{header, data[:]}, nil
}
