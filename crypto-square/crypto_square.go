package cryptosquare

import (
	"bytes"
	"math"
	"regexp"
	"strings"
)

// only ascii considered
func Encode(input string) string {
	input = normalize(input)
	row, col := getRowAndCol(len(input))

	result := &bytes.Buffer{}

	for c := 0; c < col; c++ {
		if c != 0 {
			result.WriteString(" ")
		}

		for r := 0; r < row; r++ {
			idx := r*col + c
			if idx >= len(input) {
				result.WriteString(" ")
			} else {
				result.WriteByte(input[idx])
			}
		}
	}

	return result.String()
}

var reg = regexp.MustCompile(`[^a-z1-9]`)

func normalize(input string) string {
	input = strings.ToLower(input)

	return strings.Map(func(r rune) rune {
		switch {
		case '1' <= r && r <= '9':
			fallthrough
		case 'a' <= r && r <= 'z':
			return r
		default:
			return -1
		}
	}, input)
}

// col >= sqrt(length)
// col -1 <= row <= col
func getRowAndCol(length int) (row, col int) {
	if length < 0 {
		panic("length must be greater than or equal to zero")
	}

	if length == 0 {
		return 0, 0
	}

	c := int(math.Ceil(math.Sqrt(float64(length))))
	r := (length-1)/c + 1
	return r, c
}
