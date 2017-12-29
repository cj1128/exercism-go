package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type matrix struct {
	rows [][]int
	cols [][]int
}

func New(input string) (*matrix, error) {
	rows, err := parseMatrix(input)
	if err != nil {
		return nil, err
	}

	cols := rows2cols(rows)
	return &matrix{rows, cols}, nil
}

func (m *matrix) Rows() [][]int {
	return copySlice(m.rows)
}

func (m *matrix) Cols() [][]int {
	return copySlice(m.cols)
}

func (m *matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.rows) {
		return false
	}
	if col < 0 || col >= len(m.cols) {
		return false
	}
	m.rows[row][col] = val
	m.cols[col][row] = val
	return true
}

// parseMatrix returns rows as [][]int
func parseMatrix(input string) ([][]int, error) {
	var rows [][]int
	rowNumCount := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			return nil, errors.New("empty row")
		}

		strs := strings.Fields(line)
		if rowNumCount == 0 {
			rowNumCount = len(strs)
		} else if rowNumCount != len(strs) {
			return nil, fmt.Errorf("invalid row: %s", line)
		}

		row := make([]int, rowNumCount)
		for i, str := range strs {
			n, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", str)
			}
			row[i] = n
		}

		rows = append(rows, row)
	}
	return rows, nil
}

func rows2cols(rows [][]int) [][]int {
	var cols [][]int
	if len(rows) != 0 {
		for j, _ := range rows[0] {
			var col []int
			for _, row := range rows {
				col = append(col, row[j])
			}
			cols = append(cols, col)
		}
	}
	return cols
}

func copySlice(src [][]int) [][]int {
	var result [][]int
	for _, item := range src {

		result = append(result, append([]int{}, item...))
	}
	return result
}
