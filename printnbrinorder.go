package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	s := string(n%10 + 48)
	for n >= 10 {
		n = n / 10
		s = s + string(n%10+48)
	}
	for i := '0'; i <= '9'; i++ {
		for _, num := range s {
			if i == num {
				z01.PrintRune(num)
			}
		}
	}
}
