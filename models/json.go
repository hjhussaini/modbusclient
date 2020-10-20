package models

import (
	"encoding/json"
	"io"
)

// ReadJSON deserilizes the object type from JSON string
func ReadJSON(reader io.Reader, object interface{}) error {
	return json.NewDecoder(reader).Decode(object)
}

// WriteJSON serilizes the given object into a string based JSON format
func WriteJSON(object interface{}, writer io.WriteCloser) error {
	return json.NewEncoder(writer).Encode(object)
}
