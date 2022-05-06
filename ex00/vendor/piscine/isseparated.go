package piscine

func IsSeparated(board [][]rune) bool {
	b2 := deepCopyBoard(board)
	first := true
	for y, row := range b2 {
		for x, elm := range row {
			if first && elm != 'B' {
				b2[y][x] = 'W'
				first = false
			}
			if y > 0 && b2[y-1][x] == 'W' && elm != 'B' {
				b2[y][x] = 'W'
			} else if x > 0 && b2[y][x-1] == 'W' && elm != 'B' {
				b2[y][x] = 'W'
			}
		}
		for x := len(row) - 1; x >= 0; x-- {
			elm := b2[y][x]
			if y > 0 && b2[y-1][x] == 'W' && elm != 'B' {
				b2[y][x] = 'W'
			} else if x > 0 && b2[y][x-1] == 'W' && elm != 'B' {
				b2[y][x] = 'W'
			}
		}
	}
	for _, row := range b2 {
		for _, elm := range row {
			if elm != 'W' && elm != 'B' {
				return true
			}
		}
	}
	return false
}

func deepCopyBoard(original [][]rune) [][]rune {
	copied := make([][]rune, len(original))
	for y, row := range original {
		copied[y] = make([]rune, len(row))
		for x, elm := range row {
			copied[y][x] = elm
		}
	}
	return copied
}
