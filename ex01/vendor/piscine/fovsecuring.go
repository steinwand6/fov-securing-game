package piscine

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var v_flg *int = flag.Int("d", 0, "solver visualize.")
var try int

func FOVSecuring(params []string) {
	board, ok := SettingBoard(params)
	if !ok {
		PrintStrln("Error")
		return
	}
	result, isSolved := Securing(board, 0, 0)
	if *v_flg > 0 && *v_flg < 1000 {
		fmt.Fprintf(os.Stdout, "\x1b[2J")
		fmt.Fprintf(os.Stdout, "\x1b[1;1H")
	}
	if isSolved {
		fmt.Printf("[Success(try: %d)]\n", try)
		PrintBoard(result)
	} else {
		fmt.Printf("[Error(try: %d)]\n", try)
		PrintBoard(result)
	}
}

func Securing(board [][]rune, x, y int) ([][]rune, bool) {
	m_y := len(board)
	m_x := len(board[0])
	flag.Parse()
	try++
	if *v_flg > 0 && *v_flg < 1000 {
		time.Sleep(time.Duration(*v_flg) * time.Millisecond)
		fmt.Println(v_flg)
		fmt.Fprintf(os.Stdout, "\x1b[2J")
		fmt.Fprintf(os.Stdout, "\x1b[1;1H")
		fmt.Printf("[try: %d]\n", try)
		PrintBoard(board)
	}
	if CheckFovAll(board) && !IsSeparated(board) {
		return board, true
	}
	for ; y < m_y; y++ {
		for ; x < m_x; x++ {
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
	if !ContainElementRow(board[y]) && !ContainElementCol(board, x) {
		return false
	}
	board[y][x] = 'B'
	for ch_y, row := range board {
		for ch_x := range row {
			if board[ch_y][ch_x] == '.' || board[ch_y][ch_x] == 'B' {
				continue
			}
			if ch_y < y &&
				int(board[ch_y][ch_x]-'0') < GetFovX(board, ch_x, ch_y) {
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

func ContainElementCol(board [][]rune, x int) bool {
	for _, row := range board {
		if row[x] != '.' && row[x] != 'B' {
			return true
		}
	}
	return false
}

func ContainElementRow(a []rune) bool {
	for _, v := range a {
		if v != '.' && v != 'B' {
			return true
		}
	}
	return false
}

func isOrthogonallyAdjacent(board [][]rune, x, y int) bool {
	return (y != 0 && board[y-1][x] == 'B') ||
		(x != 0 && board[y][x-1] == 'B')
}
