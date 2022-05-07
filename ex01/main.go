package main

import (
	"flag"
	"os"
	"piscine"
)

func main() {
	flag.Bool("v", false, "solver visualize")
	flag.Int("g", 5, "generate random map")
	flag.Parse()
	argc := len(os.Args)
	for i := 1; i < argc; i++ {
		f, isfile := piscine.ReadFile(os.Args[i])
		if isfile == true {
			piscine.FOVSecuringInt(f)
		} else {
			piscine.FOVSecuringInt(os.Args[1:])
			break	
		}
	}
}
