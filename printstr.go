package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	aString := []rune(str)
	for _, letter := range aString {
		z01.PrintRune(rune(letter))
	}
}
