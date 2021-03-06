package piscine

func IsSeparatedInt(b [][]int) bool {
	b2 := deepCopyBoardInt(b)
	for y, row := range b2 {
		for x, elm := range row {
			if elm != 1 {
				fillInt(b2, x, y)
				goto judge
			}
		}
	}
judge:
	for _, row := range b2 {
		for _, elm := range row {
			if !(elm == 1 || elm == -42) {
				return true
			}
		}
	}
	return false
}

func fillInt(b [][]int, x, y int) {
	if y < 0 || x < 0 || y >= len(b) || x >= len(b[0]) {
		return
	}
	if b[y][x] == 1 || b[y][x] == -42 {
		return
	}
	b[y][x] = -42
	fillInt(b, x+1, y)
	fillInt(b, x-1, y)
	fillInt(b, x, y-1)
	fillInt(b, x, y+1)
}

func deepCopyBoardInt(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for y, row := range src {
		dst[y] = make([]int, len(row))
		for x, elm := range row {
			dst[y][x] = elm
		}
	}
	return dst
}
