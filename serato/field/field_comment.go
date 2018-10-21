package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/internal/decode"
	"github.com/tombell/saga/internal/trim"
)

// Comment is the comment on the track.
type Comment struct {
	header *Header
	data   []byte
}

// Value returns the comment.
func (f *Comment) Value() string {
	s := decode.UTF16(f.data)
	return trim.Null(s)
}

func (f *Comment) String() string {
	return f.Value()
}

// NewCommentField returns a Comment, using the header to read the field data.
func NewCommentField(header *Header, r io.Reader) (*Comment, error) {
	if header.Identifier != commentID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Comment{header, data[:]}, nil
}
