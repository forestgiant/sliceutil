#sliceutil
Utility package to help with slices.

## Install
`go get -u github.com/forestgiant/sliceutil`

## Functions
### Compare
Compare will check if two slices are equal even if they aren't in the same order.
### OrderedCompare
OrderedCompare will check if two slices are equal, taking the order of the elements into consideration.
### Contains
Contains checks if a slice contains an element


## Usage
```
import "github.com/forestgiant/sliceutil"
s1 := []string{"string1", "string2"}
s2 := []string{"string1", "string2"}
element := "string2"

// Set Semantic Version
if sliceutil.Compare(s1, s2) {
  fmt.Println("Slices are equal!")
} else {
  fmt.Println("Slices aren't equal.")
}

if sliceutil.OrderedCompare(s1, s2) {
  fmt.Println("Slices are equal, considering the order of the elements!")
} else {
  fmt.Println("Slices aren't equal, considering the order of the elements.")
}

if sliceutil.Contains(s1, element) {
  fmt.Println("Slice contains the element", element)
} else {
  fmt.Println("Slice does not contain the element", element)
}
```
