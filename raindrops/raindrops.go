package raindrops

import (
	"strconv"
)

//Convert takes a number and outputs a pllngplangplong kind of string. It's complicated to explain
func Convert(num int) string {
	s := ""
	if num%3 == 0 {
		s += "Pling"
	}

	if num%5 == 0 {
		s += "Plang"
	}

	if num%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		return strconv.Itoa(num)
	}
	return s
}
