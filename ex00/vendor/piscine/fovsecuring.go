package piscine

import (
	"fmt"
)

func FOVSecuring(params []string) {
	board, ok := SettingBoard(params)
	if !ok {
		PrintStrln("Error")
	}
	fmt.Println(board)
}

func SettingBoard(strings []string) ([][]rune, bool) {
	board := make([][]rune, 0)
	for _, str := range strings {
		board = append(board, []rune(str))
	}
	if !IsValidBoard(board) {
		return [][]rune{}, false
	}
	return board, true
}

func IsValidBoard(board [][]rune) bool {
	rowCount := 0
	for _, row := range board {
		rowCount++
		// mandatory rule
		if GetLength(row) != 5 ||
			!IsValidRow(row) {
			return false
		}
	}
	// mandatory rule
	if rowCount != 5 {
		return false
	}
	return true
}

func IsValidRow(row []rune) bool {
	for _, r := range row {
		if r != '.' && !IsNumeric(r) {
			return false
		}
	}
	return true
}
