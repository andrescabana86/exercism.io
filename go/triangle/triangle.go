// Package triangle provides set of tools to get which kind of triangle it is
// based on the len of the sides
package triangle

import "math"

// Kind is type int
type Kind int

/*
NaT not a triangle
Equ equilateral
Iso isosceles
Sca scalene
**/
const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides return which kind of triangle it is
// uses the Kind type to return the value
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !(a>0) || !(b>0) || !(c>0) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	if a+b < c || b+c < a || c+a < b {
		return NaT
	}

	switch {
	case a == b && b == c && c == a:
		k = Equ
	case a == b || b == c || c == a:
		k = Iso
	case a != b && b != c && c != a:
		k = Sca
	}

	return k
}
