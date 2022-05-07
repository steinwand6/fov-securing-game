package piscine

import (
	"ft"
)

func GetLength(ary []rune) (len int) {
	for range ary {
		len++
	}
	return len
}

func IsBoardNum(r rune) bool {
	return r >= '2' && r <= '9'
}

func PrintStrln(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func PrintBoard(board [][]rune) {
	for _, row := range board {
		for _, r := range row {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}
