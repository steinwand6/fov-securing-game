package main

import (
	"flag"
	"os"
	"piscine"
)

func main() {
	g_flag := flag.Bool("g", false, "generate random maps.")
	s_flag := flag.Int("s", 5, "use with -g, size of map. -s=[2 - ]")
	f_flag := flag.Int("f", 15, "use with -g, frequency of number. -f=[2 - size * size].")
	q_flag := flag.Int("q", 2, "use with -g, generate quantity. -q=[1 - 10]")
	flag.Parse()
	if *g_flag {
		piscine.GenerateBoard(*s_flag, *f_flag, *q_flag)
		os.Exit(0)
	}
	argc := len(os.Args)
	for i := 1 + flag.NFlag(); i < argc; i++ {
		f, isfile := piscine.ReadFile(os.Args[i])
		if isfile == true {
			piscine.FOVSecuring(f)
		} else {
			piscine.FOVSecuring(os.Args[1:])
			break
		}
	}
}
