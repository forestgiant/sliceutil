package sliceutil_test

import (
	"testing"

	"github.com/forestgiant/sliceutil"
)

func TestContains(t *testing.T) {
	// Create a test struct
	type testStruct struct {
		Name string
	}

	// Used to test pointer comparison
	testStructPointer := &testStruct{"two"}

	var tests = []struct {
		s      interface{}
		e      interface{}
		result bool
	}{
		{[]string{"a", "b", "c", "d"}, "a", true},
		{[]string{"a", "b", "c", "d"}, "c", true},
		{[]string{"a", "b", "c", "d"}, "f", false},
		{[]int{1, 2, 3, 4}, 2, true},
		{[]int{1, 2, 3, 4}, 4, true},
		{[]int{1, 2, 3, 4}, 12, false},
		{[]bool{true}, true, true},
		{[]bool{false}, true, false},
		{[]bool{false}, true, false},
		{[]testStruct{{"one"}, {"two"}}, testStruct{"one"}, true},
		{[]*testStruct{{"one"}, testStructPointer}, testStructPointer, true},
	}

	for _, test := range tests {
		actual := sliceutil.Contains(test.s, test.e)
		if actual != test.result {
			t.Errorf("(%q,%q) = %v; want %v", test.s, test.e, actual, test.result)
		}
	}
}

func TestCompare(t *testing.T) {
	var tests = []struct {
		s1     interface{}
		s2     interface{}
		result bool
	}{
		{[]string{"string1", "string2"}, []string{"string1", "string2"}, true},
		{[]string{"string1", "string2"}, []string{"string2", "string1"}, true},
		{[]string{"string3", "string2"}, []string{"string2", "string1"}, false},
		{[]string{"string3", "string2", "string1"}, []string{"string2", "string1"}, false},
		{[]string{"string2", "string2", "string1"}, []string{"string2", "string1"}, false},
		{[]int{1, 2}, []int{2, 1}, true},
		{[]int{3, 2}, []int{2, 1}, false},
		{nil, nil, false},
		{nil, []int{2, 1}, false},
		{[]bool{true, false}, []bool{true, false}, true},
		{[]int{1, 0}, []bool{true, false}, false},
		{true, []bool{true, false}, false},
	}

	for _, test := range tests {
		actual := sliceutil.Compare(test.s1, test.s2)
		if actual != test.result {
			t.Errorf("(%q,%q) = %v; want %v", test.s1, test.s2, actual, test.result)
		}
	}
}

func TestOrderedCompare(t *testing.T) {
	var tests = []struct {
		s1     interface{}
		s2     interface{}
		result bool
	}{
		{[]string{"string1", "string2"}, []string{"string1", "string2"}, true},
		{[]string{"string1", "string2"}, []string{"string2", "string1"}, false},
		{[]string{"string3", "string2"}, []string{"string2", "string1"}, false},
		{[]string{"string3", "string2", "string1"}, []string{"string2", "string1"}, false},
		{[]string{"string2", "string2", "string1"}, []string{"string2", "string1"}, false},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1, 2}, []int{2, 1}, false},
		{[]int{3, 2}, []int{2, 1}, false},
		{nil, nil, true},
		{nil, []int{2, 1}, false},
		{[]bool{true, false}, []bool{true, false}, true},
		{[]int{1, 0}, []bool{true, false}, false},
		{true, []bool{true, false}, false},
		{true, true, false},
	}

	for _, test := range tests {
		actual := sliceutil.OrderedCompare(test.s1, test.s2)
		if actual != test.result {
			t.Errorf("(%q,%q) = %v; want %v", test.s1, test.s2, actual, test.result)
		}
	}
}
