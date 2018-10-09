package field

import (
	"encoding/binary"
	"io"

	"github.com/tombell/saga/strutil"
)

// Comment ...
type Comment struct {
	header *Header
	data   []byte
}

// Value ...
func (a *Comment) Value() string {
	s := strutil.DecodeUTF16(a.data)
	return strutil.TrimNull(s)
}

// NewCommentField ...
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
