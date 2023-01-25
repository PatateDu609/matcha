package payloads

import (
	"encoding/json"
	"io"
)

// Marshal writes some data to the given io.Writer.
func Marshal(data any, w io.Writer) error {
	encoder := json.NewEncoder(w)

	return encoder.Encode(data)
}
