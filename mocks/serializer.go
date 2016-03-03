package mocks

import (
	"errors"
)

// TestSerializer mock
type TestSerializer struct {
	Failure bool
}

// Serialize mock
func (s TestSerializer) Serialize(value interface{}) (ret interface{}, err error) {

	if s.Failure {
		err = errors.New("serialization error")
	} else {
		ret = "Test Value"
	}
	return
}
