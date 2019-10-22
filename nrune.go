package piscine
	func NRune(s string, n int) rune {
		var x rune = 0
		for index, char := range s {
			if index == n-1 {
				x = char
			}
		}
		return x
	}