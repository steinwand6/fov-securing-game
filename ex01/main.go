package main

import (
	"os"
	"piscine"
)

func main() {
	argc := len(os.Args)
	for i := 1; i < argc; i++ {
		f, isfile := piscine.ReadFile(os.Args[i])
		if isfile == true {
			piscine.FOVSecuring(f)
		} else {
			piscine.FOVSecuring(os.Args[1:])
			break
		}
	}
}
