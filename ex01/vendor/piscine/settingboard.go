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
	for y, row := range board {
		if y != 0 && len(row) != len(board[y-1]) ||
			!isValidRow(row) {
			return false
		}
	}
	// check square
	if len(board) == 0 || len(board) != len(board[0]) {
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
