package support

import "io"

// JSONData represents everything that can be decoded to some struct
type JSONData interface {
	DecodeJSON(r io.ReadCloser) error
}
