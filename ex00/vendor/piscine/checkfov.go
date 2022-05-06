package piscine

func GetFov(board [][]rune, x, y int) int {
	result := 1
	// ←
	for i := x - 1; i > 0; i-- {
		if board[y][i] == 'B' {
			break
		}
		result++
	}
	// →
	for i := x + 1; i < 5; i++ {
		if board[y][i] == 'B' {
			break
		}
		result++
	}
	//↑
	for i := y - 1; i > 0; i-- {
		if board[i][x] == 'B' {
			break
		}
		result++
	}
	//↓
	for i := y + 1; i < 5; i++ {
		if board[i][x] == 'B' {
			break
		}
		result++
	}
	return result
}
