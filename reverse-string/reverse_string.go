package reverse

func String(str string) string {
	runes := []rune(str)
	length := len(runes)
	result := make([]rune, length)
	for i := 0; i < len(runes); i++ {
		result[i] = runes[length-1-i]
	}
	return string(result)
}
