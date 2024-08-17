package utils

import (
	"fmt"

	"github.com/awaisamjad/db/DataType"
)

func Gap(width uint32) string {
	gap := ""
	for i := 0; i < int(width); i++ {
		gap += " "
	}
	return gap
}

func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// This returns a set so same_values isnt ordered
// TODO look at how ordering should work
func SameValues[T comparable](slice1 []T, slice2 []T) []T {
	same_values_set := make(map[T]struct{})
	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				same_values_set[slice1[i]] = struct{}{} // Add to the set
				break                                   // Avoid unnecessary comparisons once a match is found
			}
		}
	}

	// Convert set to a slice
	same_values := make([]T, 0, len(same_values_set))
	for val := range same_values_set {
		same_values = append(same_values, val)
	}

	return same_values
}

// ? Checks to see if the type for the values is at most one of DataType
func IsTypeForValuesValid[T any](datatype DataType.DataType, values []T) error {
	if len(values) == 0 {
		return fmt.Errorf("empty values slice")
	}

	var checkType func(T) bool
	var expectedType string

	switch datatype {
	case DataType.INTEGER:
		checkType = func(v T) bool { _, ok := any(v).(int); return ok }
		expectedType = "int"
	case DataType.FLOAT:
		checkType = func(v T) bool { _, ok := any(v).(float64); return ok }
		expectedType = "float64"
	case DataType.STRING:
		checkType = func(v T) bool { _, ok := any(v).(string); return ok }
		expectedType = "string"
	case DataType.CHAR:
		checkType = func(v T) bool { _, ok := any(v).(rune); return ok }
		expectedType = "string"
	default:
		return fmt.Errorf("unknown DataType")
	}

	for i, v := range values {
		if !checkType(v) {
			return fmt.Errorf("type mismatch at index %d: expected %s for %v type", i, expectedType, datatype)
		}
	}

	return nil
}

func GetValueType(value interface{}) (DataType.DataType, error) {
	switch value.(type) {
	case int:
		return DataType.INTEGER, nil
	case float64:
		return DataType.FLOAT, nil
	case string:
		return DataType.STRING, nil
	case rune:
		return DataType.CHAR, nil

	//! Default should maybe be changed to nil and then handle error
	default:
		return DataType.UNKNOWN, fmt.Errorf("type not valid")
	}
}
