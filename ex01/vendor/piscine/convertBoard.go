package piscine

func convBoardRtoInt(b [][]rune) [][]int {
	ib := make([][]int, len(b))
	for i, s := range b {
		ib[i] = make([]int, len(s))
		for j, r := range s {
			if IsNumeric(r) {
				ib[i][j] = int(r - '0')
			} else if r == '.' {
				ib[i][j] = 0
			} else if r == 'B' {
				ib[i][j] = 1
			} else {
				ib[i][j] = int(r - 'a' + 10)
			}
		}
	}
	return ib
}
