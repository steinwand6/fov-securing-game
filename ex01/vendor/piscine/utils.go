package piscine

import (
	"fmt"
	"ft"
)

func GetLength(ary []rune) (len int) {
	for range ary {
		len++
	}
	return len
}

func IsNumeric(r rune) bool {
	return r >= '0' && r <= '9'
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

func PrintBoardInt(board [][]int) {
	for _, row := range board {
		for _, r := range row {
			if r == 1 {
				fmt.Print("B ")
			} else if r < 1 {
				fmt.Print(". ")
			} else {
				c := rune(r+'0')
				fmt.Printf("%c ", c)
			}
		}
		ft.PrintRune('\n')
	}
}
