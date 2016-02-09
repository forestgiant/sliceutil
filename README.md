#sliceutil
Utility package to help with slices.

## Install
`go get -u github.com/forestgiant/sliceutil`

## Methods
### Compare
Compare will check if two slices are equal even if they aren't in the same order.
### OrderedCompare
OrderedCompare will check if two slices are equal, taking the order of the elements into consideration.

## Usage
```
import "github.com/forestgiant/sliceutil"
s1 := []string{"string1", "string2"}
s2 := []string{"string1", "string2"}

// Set Semantic Version
if Compare(s1, s2) {
  fmt.Println("Slices are equal!")
} else {
  fmt.Println("Slices aren't equal.")
}

if OrderedCompare(s1, s2) {
  fmt.Println("Slices are equal, considering the order of the elements!")
} else {
  fmt.Println("Slices aren't equal, considering the order of the elements.")
}
```
