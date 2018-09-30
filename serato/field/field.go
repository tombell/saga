package field

import "errors"

// ErrUnexpectedIdentifier is an error representing that a constructor received
// the wrong field identifier for the field type being instantiated.
var ErrUnexpectedIdentifier = errors.New("invalid field identifier")

// Field ...
type Field interface{}
