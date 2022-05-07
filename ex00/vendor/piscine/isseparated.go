package piscine

func IsSeparated(board [][]rune) bool {
	b2 := deepCopyBoard(board)
	for y, row := range b2 {
		for x, elm := range row {
			if elm != 'B' {
				fill(b2, x, y)
				goto judge
			}
		}
	}
judge:
	for _, row := range b2 {
		for _, elm := range row {
			if elm != 'W' && elm != 'B' {
				return true
			}
		}
	}
	return false
}

func fill(board [][]rune, x, y int) {
	if y < 0 || x < 0 || y >= len(board) || x >= len(board[0]) {
		return
	}
	if board[y][x] == 'W' || board[y][x] == 'B' {
		return
	}
	board[y][x] = 'W'
	fill(board, x+1, y)
	fill(board, x-1, y)
	fill(board, x, y-1)
	fill(board, x, y+1)
}

func deepCopyBoard(original [][]rune) [][]rune {
	copied := make([][]rune, len(original))
	for y, row := range original {
		copied[y] = make([]rune, len(row))
		for x, elm := range row {
			copied[y][x] = elm
		}
	}
	return copied
}
