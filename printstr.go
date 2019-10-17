package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	slice := []rune(str)
	for _, a := range slice {
		z01.PrintRune(rune(a))
	}
	z01.PrintRune(10)
}
