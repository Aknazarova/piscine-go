package main

import "github.com/01-edu/z01"
import "os"

func main() {
	str := os.Args
	count := 0
	for range str {
		count++
	}
	for i := 1; i < count - 1; i++ { // ne scitat' 1-yi symbol
		for _, char := range str[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune(10)
	}
}
