package main

import (
	"os"
	"piscine"
)

func main() {
	params := os.Args[1:]
	piscine.FOVSecuring(params)
}
