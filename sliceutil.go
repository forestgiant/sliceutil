package sliceutil

import "reflect"

// Compare will check if two slices are equal
// even if they aren't in the same order
// Inspired by github.com/stephanbaker white board sudo code
func Compare(s1, s2 interface{}) bool {
	if s1 == nil || s2 == nil {
		return false
	}

	// Convert slices to correct type
	slice1 := convertSliceToInterface(s1)
	slice2 := convertSliceToInterface(s2)
	if slice1 == nil || slice2 == nil {
		return false
	}

	if len(slice1) != len(slice2) {
		return false
	}

	// setup maps to store values and count of slices
	m1 := make(map[interface{}]int)
	m2 := make(map[interface{}]int)

	for i := 0; i < len(slice1); i++ {
		// Add each value to map and increment for each found
		m1[slice1[i]]++
		m2[slice2[i]]++
	}

	for key := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}

	return true
}

// convertSliceToInterface takes a slice passed in as an interface{}
// then converts the slice to a slice of interfaces
func convertSliceToInterface(s interface{}) (slice []interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		return nil
	}

	length := v.Len()
	slice = make([]interface{}, length)
	for i := 0; i < length; i++ {
		slice[i] = v.Index(i).Interface()
	}

	return slice
}
