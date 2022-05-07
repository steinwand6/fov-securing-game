package piscine

import "fmt"

func FOVSecuringInt(params []string) {
	fmt.Println(params)
	brd, ok := SettingBoard(params)
	if !ok {
		PrintStrln("Error")
		return
	}
	ib := convBoardRtoInt(brd)
	initBoard(ib)
	fmt.Println(ib)
	// result, isSolved := Securing(ib, 0, 0)
	// if isSolved {
	// 	PrintBoard(result)
	// } else {
	// 	PrintStrln("Error")
	// }
}

func initBoard(brd [][]int) {
	max := len(brd) + len(brd[0]) - 1
	max /= 2
	for i, iy := range brd {
		for j, ix := range iy {
			if ix == 0 && !(isElemInCol(brd, j) || isElemInRow(brd[i])) {
				brd[i][j] = -1
			} else if ix > max {
				placeUnplaceable(brd, j, i)
			}
		}
	}
}

func placeUnplaceable(brd [][]int, x, y int) {
	// 現状 B を考慮していない
	left := x
	right := len(brd[0]) - x - 1
	top := y
	bottom := len(brd) - y - 1
	max := len(brd) + len(brd[0]) - 1
	// left
	l := left - max + brd[y][x]
	for i := 0; i < l; i++ {
		if brd[y][x-i] == 0 {
			brd[y][x-i] = -1
		}
	}
	// right
	r := right - max + brd[y][x]
	for i := 0; i < r; i++ {
		if brd[y][x+i] == 0 {
			brd[y][x+i] = -1
		}
	}
	// top
	t := top - max + brd[y][x]
	for i := 0; i < t; i++ {
		if brd[y-i][x] == 0 {
			brd[y-i][x] = -1
		}
	}
	// bottom
	b := bottom - max + brd[y][x]
	for i := 0; i < b; i++ {
		if brd[y+i][x] == 0 {
			brd[y+i][x] = -1
		}
	}
}

// func Securing(brd [][]int, x, y int) ([][]int, bool) {
// 	size_x := len(brd[0])
// 	size_y := len(brd)
// 	if CheckFovAll(brd) && !IsSeparated(brd) {
// 		return brd, true
// 	}
// 	for ; y < size_y; y++ {
// 		for ; x < size_x; x++ {
// 			if placeBlackSquare(brd, x, y) {
// 				result, comp := Securing(brd, x, y)
// 				if comp {
// 					return result, comp
// 				}
// 				brd[y][x] = -1
// 			}
// 		}
// 		x = 0
// 	}
// 	return brd, false
// }

// func placeBlackSquare(brd [][]int, x, y int) bool {
// 	if brd[y][x] != '.' {
// 		return false
// 	}
// 	if isOrthogonallyAdjacent(brd, x, y) {
// 		return false
// 	}
// 	if !ContainElementRow(brd[y]) && !ContainElementCol(brd, x) {
// 		return false
// 	}
// 	brd[y][x] = 'B'
// 	for ch_y, row := range brd {
// 		for ch_x := range row {
// 			if brd[ch_y][ch_x] == '.' || brd[ch_y][ch_x] == 'B' {
// 				continue
// 			}
// 			if ch_y < y &&
// 				int(brd[ch_y][ch_x]-'0') < GetFovX(brd, ch_x, ch_y) {
// 				brd[y][x] = '.'
// 				return false
// 			}
// 			if int(brd[ch_y][ch_x]-'0') > GetFov(brd, ch_x, ch_y) {
// 				brd[y][x] = '.'
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func isElemInCol(brd [][]int, x int) bool {
	for _, row := range brd {
		if row[x] != 0 && row[x] != 1 {
			return true
		}
	}
	return false
}

func isElemInRow(a []int) bool {
	for _, v := range a {
		if v != 0 && v != 1 {
			return true
		}
	}
	return false
}

// func isOrthogonallyAdjacent(brd [][]int, x, y int) bool {
// 	return (y != 0 && brd[y-1][x] == 1) ||
// 		(x != 0 && brd[y][x-1] == 1)
// }
