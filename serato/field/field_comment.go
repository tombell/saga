package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Comment is the comment on the track.
type Comment struct {
	header *Header
	data   []byte
}

// Value returns the comment.
func (c *Comment) Value() string {
	s := strutil.DecodeUTF16(c.data)
	return strutil.TrimNull(s)
}

func (c *Comment) String() string {
	return c.Value()
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
