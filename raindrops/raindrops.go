package raindrops

import (
	"bytes"
	"math"
	"strconv"
)

//Convert takes a number and outputs a pllngplangplong kind of string. It's complicated to explain
func Convert(num int) string {
	var b bytes.Buffer
	for _, n := range factors(num) {
		if n == 3 {
			b.WriteString("Pling")
		}

		if n == 5 {
			b.WriteString("Plang")
		}

		if n == 7 {
			b.WriteString("Plong")
		}
	}

	s := b.String()
	if s == "" {
		return strconv.FormatInt(int64(num), 10)
	}
	return s
}

func factors(num int) []int {
	var a []int
	max := int(math.Max(8, math.Sqrt(float64(num))))
	for i := 1; i < max; i++ {
		if (i == 3 || i == 5 || i == 7) && num%i == 0 {
			a = append(a, i)
		}
	}
	return a
}
