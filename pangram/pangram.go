package pangram

import "strings"

const testVersion = 2

func IsPangram(str string) bool {
	str = strings.ToLower(str)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(str, i) {
			return false
		}
	}
	return true
}
