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

func IsNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

func PrintStrln(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
