package grains

import "errors"

const testVersion = 1

func Square(n int) (uint64, error) {
	if n > 64 || n <= 0 {
		return 0, errors.New("bad input")
	}

	return 1 << (uint(n) - 1), nil
}

// value is 2^64 - 1
func Total() uint64 {
	return ^uint64(0)
}
