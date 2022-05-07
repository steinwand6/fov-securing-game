package piscine

func CheckFovAll(board [][]rune) bool {
	for y, row := range board {
		for x, elm := range row {
			if elm == '.' || elm == 'B' {
				continue
			}
			if !CheckFov(board, x, y) {
				return false
			}
		}
	}
	return true
}

func CheckFov(board [][]rune, x, y int) bool {
	return int(board[y][x]-'0') == GetFov(board, x, y)
}

func GetFov(board [][]rune, x, y int) int {
	return GetFovX(board, x, y) + GetFovY(board, x, y) - 1
}

func GetFovX(board [][]rune, x, y int) int {
	m := len(board[0])
	result := 1
	// ←
	for i := x - 1; i >= 0; i-- {
		if board[y][i] == 'B' {
			break
		}
		result++
	}
	// →
	for i := x + 1; i < m; i++ {
		if board[y][i] == 'B' {
			break
		}
		result++
	}
	return result
}

func GetFovY(board [][]rune, x, y int) int {
	m := len(board)
	result := 1
	//↑
	for i := y - 1; i >= 0; i-- {
		if board[i][x] == 'B' {
			break
		}
		result++
	}
	//↓
	for i := y + 1; i < m; i++ {
		if board[i][x] == 'B' {
			break
		}
		result++
	}
	return result
}
