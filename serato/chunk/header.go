package chunk

// Header ...
type Header struct {
	Identifier [4]byte
	Length     uint32
}

// Type ...
func (h *Header) Type() string {
	return string(h.Identifier[:])
}

// NewHeader ...
func NewHeader() *Header {
	return &Header{}
}
