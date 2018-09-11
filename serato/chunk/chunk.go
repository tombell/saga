package chunk

// Chunk ...
type Chunk interface {
	Header() *Header
	Type() string
}
