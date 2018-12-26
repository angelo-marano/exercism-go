package hamming

import "fmt"

//Distance calculates the hamming distance of two string, returns an error if the strings are of different length
func Distance(a, b string) (int, error) {
	var count int
	if len(a) != len(b) {
		return 0, fmt.Errorf("unequal lengths %q, %q", a, b)
	}
	for i, c := range a {
		if c != rune(b[i]) {
			count++
		}
	}

	return count, nil
}
