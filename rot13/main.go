package main

import (
	"github.com/01-edu/z01"
	"os"
)

func Rot13(str string) string {
	res := []rune(str)
	for i, val := range str {
		if !(('A' <= val && val <= 'Z') || ('a' <= val && val <= 'z')) {
			continue
		}
		cap := rune(0)
		if 'A' <= val && val <= 'Z' {
			cap = 'a' - 'A'
		}
		if val+cap+13 > 'z' {
			res[i] = 'a' + (val + cap + 13 - 'z') - cap - 1
			continue
		}
		res[i] = val + 13
	}
	return string(res)
}

func main() {
	if len(os.Args) == 2 {
		res := os.Args[1]
		runes := []rune(Rot13(res))
		for i := range runes {
			z01.PrintRune(runes[i])
		}
	}
}
