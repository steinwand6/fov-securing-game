package piscine

import "fmt"

func FOVSecuringInt(params []string) {
	brd, ok := SettingBoard(params)
	if !ok {
		PrintStrln("map Error")
		return
	}
	ib := convBoardRtoInt(brd)
	initBoard(ib)
	fmt.Println(ib)
	result, isSolved := SecuringInt(ib, 0, 0)
	if isSolved {
		PrintBoardInt(result)
	} else {
		PrintStrln("resolve Error")
	}
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

func isElemInCol(brd [][]int, x int) bool {
	for _, row := range brd {
		if row[x] >= 2 {
			return true
		}
	}
	return false
}

func isElemInRow(a []int) bool {
	for _, v := range a {
		if v >= 2 {
			return true
		}
	}
	return false
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

func SecuringInt(brd [][]int, x, y int) ([][]int, bool) {
	size_x := len(brd[0])
	size_y := len(brd)
	if CheckFovAllInt(brd) && !IsSeparatedInt(brd) {
		return brd, true
	}
	for ; y < size_y; y++ {
		for ; x < size_x; x++ {
			if placeBlackInt(brd, x, y) {
				result, comp := SecuringInt(brd, x, y)
				if comp {
					return result, comp
				}
				brd[y][x] = 0
			}
		}
		x = 0
	}
	return brd, false
}

func placeBlackInt(brd [][]int, x, y int) bool {
	if brd[y][x] != 0 || isClosedBlack(brd, x, y) {
		return false
	}
	brd[y][x] = 1
	for iy, row := range brd {
		for ix := range row {
			if brd[iy][ix] <= 1 {
				continue
			}
			if iy < y && brd[iy][ix] < GetFovXInt(brd, ix, iy) {
				brd[y][x] = 0
				return false
			}
			if brd[iy][ix] > GetFovInt(brd, ix, iy) {
				brd[y][x] = 0
				return false
			}
		}
	}
	return true
}


// origin : orthogonallyAdjascent
func isClosedBlack(brd [][]int, x, y int) bool {
	return (y != 0 && brd[y-1][x] == 1) ||
		(x != 0 && brd[y][x-1] == 1)
}
