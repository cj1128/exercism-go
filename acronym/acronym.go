package acronym

import "strings"

const testVersion = 3

func Abbreviate(str string) string {
	str = strings.Replace(str, "-", " ", -1)
	result := ""
	for _, item := range strings.Fields(str) {
		first := string([]rune(item)[0])
		result += strings.ToUpper(first)
	}
	return result
}
