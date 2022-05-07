package piscine

import "fmt"

func FOVSecuring(params []string) {
	board, ok := SettingBoard(params)
	if !ok {
		PrintStrln("Error")
		return
	}
	result, isSolved := Securing(board, 0, 0)
	if isSolved {
		PrintBoard(result)
	} else {
		PrintStrln("Error")
	}
}

func Securing(board [][]rune, x, y int) ([][]rune, bool) {
	if CheckFovAll(board) && !IsSeparated(board) {
		return board, true
	}
	for ; y < 5; y++ {
		for ; x < 5; x++ {
			if placeBlackSquare(board, x, y) {
				result, comp := Securing(board, x, y)
				if comp {
					return result, comp
				}
				board[y][x] = '.'
			}
		}
		x = 0
	}
	return board, false
}

func placeBlackSquare(board [][]rune, x, y int) bool {
	if board[y][x] != '.' {
		return false
	}
	if isOrthogonallyAdjacent(board, x, y) {
		return false
	}
	board[y][x] = 'B'
	for ch_y, row := range board {
		for ch_x := range row {
			if board[ch_y][ch_x] == '.' || board[ch_y][ch_x] == 'B' {
				continue
			}
			if ch_y < y && int(board[ch_y][ch_x]-'0') < GetFovX(board, x, y) {
				fmt.Println(int(board[ch_y][ch_x]-'0'), GetFovX(board, x, y))
				board[y][x] = '.'
				return false
			}
			if int(board[ch_y][ch_x]-'0') > GetFov(board, ch_x, ch_y) {
				board[y][x] = '.'
				return false
			}
		}
	}
	return true
}

func isOrthogonallyAdjacent(board [][]rune, x, y int) bool {
	return (y != 0 && board[y-1][x] == 'B') ||
		(x != 0 && board[y][x-1] == 'B')
}
