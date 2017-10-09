package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("bad arguments")
	}

	result := 0
	as := []rune(a)
	bs := []rune(b)
	for i, c := range as {
		if bs[i] != c {
			result++
		}
	}
	return result, nil
}
