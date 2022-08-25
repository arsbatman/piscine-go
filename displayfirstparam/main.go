package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args[1]
	for _, letters := range arguments {
		z01.PrintRune(letters)
	}
}
