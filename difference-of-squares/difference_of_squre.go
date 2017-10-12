package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
