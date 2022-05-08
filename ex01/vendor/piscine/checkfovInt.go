package piscine

func CheckFovAllInt(b [][]int) bool {
	for y, row := range b {
		for x, elm := range row {
			if elm <= 1 {
				continue
			} else if !CheckFov(b, x, y){
				return false
			}
		}
	}
	return true
}

func CheckFov(b [][]int, x, y int) bool {
	return b[y][x] == GetFov(b, x, y) // return int is better ?? pending
}

func GetFovInt(b [][]int, x, y int) int {
	return GetFovX(b, x, y) + GetFovY(b, x, y) - 1
}

func GetFovX(b [][]int, x, y int) int {
	m := len(b[0])
	fov := 1
	// left
	for i := x - 1; i >= 0; i-- {
		if b[y][i] == 1 {
			break
		}
		fov++
	}
	// right
	for i := x + 1; i < m; i++ {
		if b[y][i] == 1 {
			break
		}
		fov++
	}
	return fov
}

func GetFovY(b [][]int, x, y int) int {
	m := len(b)
	fov := 1
	// top
	for i := y - 1; i >= 0; i-- {
		if b[i][x] == 1 {
			break
		}
		fov++
	}
	// bottom
	for i := y + 1; i < m; i++ {
		if b[i][x] == 1 {
			break
		}
		fov++
	}
	return fov
}
