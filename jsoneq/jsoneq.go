package jsoneq

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ConvertByte converts a byte array to an object which can be checked if it is equal to the other value as JSON.
func ConvertByte(b []byte) (interface{}, error) {
	var d interface{}
	err := json.Unmarshal(b, &d)
	if err == nil {
		return d, nil
	}
	return nil, fmt.Errorf("failed to unmarshal json: %w", err)
}

// Convert converts a given value to an object which can be checked if it is equal to the other value as JSON.
func Convert(x interface{}) (interface{}, error) {
	if a, ok := x.([]byte); ok {
		return ConvertByte(a)
	}
	b, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("faled to marshal json: %w", err)
	}
	return ConvertByte(b)
}

// Equal checks if two values are equal as JSON.
func Equal(x, y interface{}) (bool, error) {
	if reflect.DeepEqual(x, y) {
		return true, nil
	}
	a, err := Convert(x)
	if err != nil {
		return false, err
	}
	b, err := Convert(y)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(a, b), nil
}
