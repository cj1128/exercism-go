package raindrops

import "strconv"

const testVersion = 3

func Convert(n int) string {
	result := ""

	if n%3 == 0 {
		result += "Pling"
	}

	if n%5 == 0 {
		result += "Plang"
	}

	if n%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(n)
	}

	return result
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
