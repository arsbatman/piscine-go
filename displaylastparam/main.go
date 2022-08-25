package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	size := len(os.Args)
	arguments := os.Args[size-1]
	if size > 1 {
		for _, letters := range arguments {
			z01.PrintRune(letters)
		}
	}
}
