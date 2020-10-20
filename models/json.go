package models

import (
	"encoding/json"
	"io"
)

// ReadJSON deserilizes the object type from JSON string
func ReadJSON(reader io.Reader, object interface{}) error {
	return json.NewDecoder(reader).Decode(object)
}
