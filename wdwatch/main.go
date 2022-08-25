package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		word1 := os.Args[1]
		word2 := os.Args[2]
		rune1 := []rune(word1)
		rune2 := []rune(word2)
		count := 0
		for i := 0; i < len(rune1); i++ {
			for j := 0; j < len(rune2); j++ {
				if rune1[i] == rune2[j] {
					count++
					i++
				}
			}
		}
		if count == len(rune1) {
			for _, ch := range rune1 {
				z01.PrintRune(ch)
			}
		}
	}

}
