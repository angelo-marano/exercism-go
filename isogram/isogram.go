package isogram

import (
	"sort"
	"strings"
	"unicode"
)

//IsIsogram determines whether or not a word is an isogram (word or phrase without a repeating letter)
func IsIsogram(s string) bool {

	runes := []rune(strings.ToLower(s))
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for i, n := range runes {
		if i == len(runes)-2 {
			return true
		}
		if unicode.IsLetter(n) && n == runes[i+1] {
			return false
		}
	}

	return true
}
