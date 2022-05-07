package piscine

import "os"

func ReadFile(file string) ([]string, bool) {
	s, isfile := Read(file)
	return FormatFile(s), isfile
}

func Read(file string) (string, bool) {
	f, err := os.Open(file)
	if err != nil {
		return "", false
	}
	defer f.Close()
	buf := make([]byte, 1024)
	str := ""
	for {
		n, err := f.Read(buf)
		if n == 0 || err != nil {
			break
		}
		str += string(buf[:n])
	}
	if err != nil {
		return "", false
	}
	return str, true
}

func FormatFile(s string) []string {
	row := CountRows(s)
	ret := make([]string, 0)
	idx := 0
	for i, c := range s {
		if c == '\n' {
			ret = append(ret, s[idx:i])
			idx = i+1
			row--
		}
	}
	if row > 0 {
		ret = append(ret, s[idx:])
	}
	return ret
}

func CountRows(s string) int {
	row := 1
	for _, c := range s {
		if c == '\n' {
			row++
		}
	}
	return row
}