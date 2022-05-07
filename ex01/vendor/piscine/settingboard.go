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
	col := 0
	row := 0
	if len(board) == 0 {
		return false
	}
	for i, ir := range board {
		if i == 0 {
			col = len(ir)
		} else if len(ir) != col {
			return false
		}
		row++
	}
	if !isValidChar(board, col, row) {
		return false
	}
	return true
}

func isValidChar(b [][]rune, col, row int) bool {
	max := col + row - 1
	max_c := '0'
	over_ten := false
	if max >= 10 {
		max_c = rune(int('a') + (max - 10))
		over_ten = true
	} else {
		max_c = rune(int('0') + max)
	}
	for _, r := range b {
		for _, c := range r {
			if over_ten {
				if !((c >= '2' && c <= '9') || (c >= 'a' && c <= max_c) || c == '.' || c == 'B') {
					return false
				}
			} else {
				if !((c >= '2' && c <= max_c) || c == '.' || c == 'B') {
					return false
				}
			}
		}
	}
	return true
}
