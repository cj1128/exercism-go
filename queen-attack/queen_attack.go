package queenattack

import "errors"

const testVersion = 2

var errSameSquare = errors.New("same square")
var errInvalidPosition = errors.New("invalid position")

func CanQueenAttack(p1, p2 string) (bool, error) {
	if p1 == p2 {
		return false, errSameSquare
	}
	row1, col1, err := parsePosition(p1)
	if err != nil {
		return false, err
	}

	row2, col2, err := parsePosition(p2)
	if err != nil {
		return false, err
	}

	if row1 == row2 || col1 == col2 || isDiagonal(row1, col1, row2, col2) {
		return true, nil
	}

	return false, nil
}

func isDiagonal(row1, col1, row2, col2 int) bool {
	d := (row1 - row2) / (col1 - col2)
	return d == 1 || d == -1
}

// return (row, col, error), row is 'a' ~ 'h', col is '1' ~ '8'
func parsePosition(p string) (int, int, error) {
	if len(p) != 2 {
		return 0, 0, errInvalidPosition
	}

	row := int(p[0])
	col := int(p[1])

	if row >= 'a' && row <= 'h' && col >= '1' && col <= '8' {
		return row, col, nil
	}
	return 0, 0, errInvalidPosition
}
