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
	row := 0
	for _, ir := range board {
		if len(ir) != 5 || !isValidRow(ir){
			return false
		}
		row++
	}
	// mandatory rule
	if row != 5 {
		return false
	}
	return true
}

func isValidRow(row []rune) bool {
	for _, r := range row {
		if r != '.' && !IsBoardNum(r) {
			return false
		}
	}
	return true
}
