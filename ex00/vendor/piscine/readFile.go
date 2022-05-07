package piscine

import "os"

func ReadFile(file string) []string {
	return (FormatFile(Read(file)))
}

func Read(file string) string {
	f, err := os.Open(file)
	if err != nil {
		PrintStrln(err.Error())
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
	return str
}

func FormatFile(s string) []string {
	row := CountRows(s)
	ret := make([]string, 0)
	idx := 0
	for i, c := range s {
		if c == '\n' {
			ret = append(ret, s[idx:i+1])
			idx = i+1
			row--
		}
	}
	if row > 0 {
		ret = appned(ret, s[idx:])
	}
	return ret
}

func CountRows(s stirng) int {
	row := 1
	for _, c := range s {
		if c == '\n' {
			row++
		}
	}
	return row
}