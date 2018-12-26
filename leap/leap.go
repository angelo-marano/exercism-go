package leap

// IsLeapYear determines if a year is a leap year
func IsLeapYear(i int) bool {
	return ((i%4 == 0 && i%100 != 0) || i%400 == 0)
}
