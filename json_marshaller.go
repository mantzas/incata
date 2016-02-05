package incata

import (
	"encoding/json"
)

// JSONMarshaller JSON Serializer
type JSONMarshaller struct {
}

// NewJSONMarshaller creates a new JSON serializer
func NewJSONMarshaller() *JSONMarshaller {

	return new(JSONMarshaller)
}

// Serialize JSON Implementation of serialization
func (s *JSONMarshaller) Serialize(value interface{}) (interface{}, error) {

	jsonArray, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(jsonArray), nil
}

// Deserialize JSON implementation of deserialization
func (s *JSONMarshaller) Deserialize(value interface{}, result interface{}) error {

	if err := json.Unmarshal([]byte(value.(string)), result); err != nil {
		return err
	}

	return nil
}
