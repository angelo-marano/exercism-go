package hamming

import (
	"errors"
)

//Distance calculates the hamming distance of two string, returns an error if the strings are of different length
func Distance(a, b string) (int, error) {
	var aLen = len(a)
	var count int
	var e error
	if aLen != len(b) {
		e = errors.New("The string lengths do not match")
	} else {
		for i, c := range a {
			if c != rune(b[i]) {
				count++
			}
		}
	}

	return count, e
}
