// Package leap implements IsLeapYear function
package leap

// IsLeapYear returns true if year is leap
func IsLeapYear(year int) bool {
	if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
		return false
	}
	return true
}
