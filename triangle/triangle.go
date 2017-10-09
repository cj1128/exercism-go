package triangle

import "math"

const testVersion = 3

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = iota // not a triangle
	Equ      = iota // equilateral
	Iso      = iota // isosceles
	Sca      = iota // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if !validSide(a) || !validSide(b) || !validSide(c) {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && a == c && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

func validSide(n float64) bool {
	return n > 0 && !math.IsInf(n, 1)
}
