package secret

const testVersion = 2

func Handshake(code uint) []string {
	var result []string

	if isBitSet(code, 0) {
		result = append(result, "wink")
	}

	if isBitSet(code, 1) {
		result = append(result, "double blink")
	}

	if isBitSet(code, 2) {
		result = append(result, "close your eyes")
	}

	if isBitSet(code, 3) {
		result = append(result, "jump")
	}

	if isBitSet(code, 4) {
		result = reverseStrings(result)
	}

	return result
}

func isBitSet(n uint, bit uint) bool {
	return n&(1<<bit) != 0
}

func reverseStrings(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
