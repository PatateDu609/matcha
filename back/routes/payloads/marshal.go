package payloads

import (
	"encoding/json"
	"io"
)

// Marshal takes a request's body as parameter and marshals it to return a type-safe object containing the request's data.
func Marshal[T any](input io.ReadCloser) (T, error) {
	var payload T
	decoder := json.NewDecoder(input)

	if err := decoder.Decode(&payload); err != nil {
		return payload, err
	}
	return payload, nil
}
