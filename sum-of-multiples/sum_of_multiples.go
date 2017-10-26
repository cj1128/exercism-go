package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if dividible(i, divisors) {
			sum += i
		}
	}
	return sum
}

func dividible(n int, divisors []int) bool {
	for _, d := range divisors {
		if n%d == 0 {
			return true
		}
	}
	return false
}
