package utils

import (
	"github.com/awaisamjad/db/Type"
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

func CheckValueType(value interface{}) Type.Type {
	switch value.(type) {
	case int:
		return Type.INTEGER
	case float64:
		return Type.FLOAT
	case string:
		return Type.STRING
	case rune:
		return Type.CHAR

	//! Default should maybe be changed to nil and then handle error
	default:
		return ""
	}
}
