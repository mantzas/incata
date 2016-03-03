package marshal

import (
	"encoding/json"
	"errors"
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

	data, found := value.(string)

	if !found {
		return errors.New("value is not a string")
	}

	if err := json.Unmarshal([]byte(data), result); err != nil {
		return err
	}

	return nil
}
