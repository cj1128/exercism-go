package isogram

import "strings"

func IsIsogram(str string) bool {
	// only consider ascii
	m := make(map[rune]bool, len(str))
	for _, r := range strings.ToLower(str) {
		if r < 'a' || r > 'z' {
			continue
		}
		if m[r] == true {
			return false
		}
		m[r] = true
	}
	return true
}
