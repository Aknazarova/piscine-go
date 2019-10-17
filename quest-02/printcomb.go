package piscine

import "github.com/01-edu/z01"

//PrintRune PrintComb
func PrintComb() {
	for x := '0'; x <= '9'; x++ {
		for y := x + 1; y <= '9'; y++ {
			for z := y + 1; z <= '9'; z++ {
				if x < y && y < z {
					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(z)
					if x != '7' || y != '8' || z != '9' {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
