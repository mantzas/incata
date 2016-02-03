package incata

import (
	"encoding/json"
)

// JSONSerializer JSON Serializer
type JSONSerializer struct {
}

// NewJSONSerializer creates a new JSON serializer
func NewJSONSerializer() *JSONSerializer {

	return new(JSONSerializer)
}

// Serialize JSON Implementation of serialization
func (serializer *JSONSerializer) Serialize(value interface{}) (interface{}, error) {

	jsonArray, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}
	
	return string(jsonArray), nil
}
