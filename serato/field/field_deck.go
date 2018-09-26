package field

import (
	"encoding/binary"
	"io"
)

// TODO: Deck is field #31

// Deck ...
type Deck struct {
	header *Header
	data   []byte
}

// Value ...
func (d *Deck) Value() int {
	return int(binary.BigEndian.Uint32(d.data))
}

// NewDeckField ...
func NewDeckField(header *Header, r io.Reader) (*Deck, error) {
	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Deck{header, data[:]}, nil
}
