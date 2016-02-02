package incata

import (
	"encoding/json"
)

// JSONSerializer JSON Serializer
type JSONSerializer struct {
}

// NewJSONSerializer creates a new JSON serializer
func NewJSONSerializer() *JSONSerializer {

	return &JSONSerializer{}
}

// Serialize JSON Implementation of serialization
func (serializer *JSONSerializer) Serialize(value interface{}) (ret string, err error) {

	jsonArray, err := json.Marshal(value)

	if err != nil {
		return
	}

	ret = string(jsonArray)
	return
}
