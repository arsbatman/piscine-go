package main

import (
	"os"
)

func main() {
	if len(os.Args) == 2 {
		counter := 0
		res := 0
		rune1 := os.Args[1]
		runes := []rune(rune1)
		for i := 0; i < len(runes); i++ {
			res = int(runes[i] - '0')
		}
		for j := 1; j <= res; j *= 2 {
			counter = j
		}
		check(counter, res)
	}
}

func check(a1, b1 int) bool {
	if a1 == b1 {
		return true
	} else {
		return false
	}
}
