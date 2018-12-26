package triangle

import (
	"sort"
)

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene
)

func checkIfTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	return sides[0]+sides[1] >= sides[2]
}

// KindFromSides returns the type of triangle
func KindFromSides(a, b, c float64) Kind {

	if !checkIfTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
