package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) []Triplet {
	var result []Triplet
	for a := min; a <= max-2; a++ {
	START:
		for b := a + 1; b <= max-1; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					result = append(result, Triplet{a, b, c})
					break START
				}
			}
		}
	}
	return result
}

func Sum(p int) []Triplet {
	var result []Triplet
	for a := 1; a < p/3; a++ {
		for b := a + 1; b < p/2; b++ {
			c := p - a - b
			if c < b {
				break
			}
			if a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return result
}
