// Package leap provides func to know if year is leap or not
package leap

// IsLeapYear should return true if year param is true
// otherwise returns false
func IsLeapYear(year int) bool {
	if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
		return false
	}
	return true
}
