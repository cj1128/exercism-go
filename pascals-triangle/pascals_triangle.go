package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, i+1)
		row := result[i]

		for j := range row {
			if j == 0 || j == len(row)-1 {
				row[j] = 1
			} else {
				parentRow := result[i-1]
				row[j] = parentRow[j-1] + parentRow[j]
			}
		}
	}

	return result
}
