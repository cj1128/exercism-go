package series

const testVersion = 2

func All(n int, s string) []string {
	result := make([]string, len(s)-n+1)
	for i := range result {
		result[i] = s[i : i+n]
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
