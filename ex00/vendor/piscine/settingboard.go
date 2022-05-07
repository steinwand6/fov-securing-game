package piscine

func SettingBoard(strings []string) ([][]rune, bool) {
	board := make([][]rune, 0)
	if strings == nil {
		return board, false
	}
	for _, str := range strings {
		board = append(board, []rune(str))
	}
	if !isValidBoard(board) {
		return [][]rune{}, false
	}
	return board, true
}

func isValidBoard(board [][]rune) bool {
	rowCount := 0
	for _, row := range board {
		rowCount++
		// mandatory rule
		if GetLength(row) != 5 ||
			!isValidRow(row) {
			return false
		}
	}
	// mandatory rule
	if rowCount != 5 {
		return false
	}
	return true
}

func isValidRow(row []rune) bool {
	for _, r := range row {
		if r != '.' && !IsNumeric(r) {
			return false
		}
	}
	return true
}
