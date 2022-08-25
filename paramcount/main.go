package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	c := '0'
	words := os.Args[1:]
	for range words {
		c++
	}
	z01.PrintRune(c)
}
