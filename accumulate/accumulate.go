package accumulate

const testVersion = 1

func Accumulate(items []string, fun func(string) string) []string {
	var result []string
	for _, item := range items {
		result = append(result, fun(item))
	}
	return result
}
