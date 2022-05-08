package main

import (
	"flag"
	"os"
	"piscine"
)

var (
	v_flg bool
	g_flg int
)

func init() {
	//flag.BoolVar(&v_flg, "v", false, "solver visualize")
	flag.IntVar(&g_flg, "g", 5, "generate random map")
}

func main() {
	flag.Parse()
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
