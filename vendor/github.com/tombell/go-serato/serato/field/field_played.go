package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Played is the played status of the track in Serato.
type Played struct {
	header *Header
	data   []byte
}

// Value returns the played status.
func (f *Played) Value() bool {
	return f.data[0] == byte(1)
}

func (f *Played) String() string {
	return fmt.Sprintf("%v", f.Value())
}

// NewPlayedField returns a Played, using the header to read the field data.
func NewPlayedField(header *Header, r io.Reader) (*Played, error) {
	if header.Identifier != playedID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Played{
		header: header,
		data:   data[:],
	}, nil
}
