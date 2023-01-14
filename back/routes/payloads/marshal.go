package payloads

import (
	"encoding/json"
	"io"
)

// Marshal takes a request's struct as parameter and marshals it to return a type-safe object containing the request's body.
func Marshal(data any, w io.Writer) error {
	encoder := json.NewEncoder(w)

	return encoder.Encode(data)
}
