package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}
