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
	fmt.Print("\033[7m")
	for _, row := range board {
		for _, r := range row {
			switch r {
			case 'B':
				fmt.Print("\033[37m")
				fmt.Print("\033[40m")
				fmt.Print("B")
			default:
				fmt.Print("\033[39m")
				fmt.Print("\033[49m")
				ft.PrintRune(r)
			}
			ft.PrintRune(' ')
		}
		ft.PrintRune('\n')
	}
	fmt.Print("\033[39m")
	fmt.Print("\033[49m")
}

func PrintBoardInt(board [][]int) {
	fmt.Print("\033[7m")
	for _, row := range board {
		for _, r := range row {
			if r == 1 {
				fmt.Print("\033[37m")
				fmt.Print("\033[40m")
				fmt.Print("B ")
			} else if r < 1 {
				fmt.Print("\033[39m")
				fmt.Print("\033[49m")
				fmt.Print(". ")
			} else {
				fmt.Print("\033[39m")
				fmt.Print("\033[49m")
				c := rune(r + '0')
				fmt.Printf("%c ", c)
			}
		}
		ft.PrintRune('\n')
	}
	fmt.Print("\033[39m")
	fmt.Print("\033[49m")
}
