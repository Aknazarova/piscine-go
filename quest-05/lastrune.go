package piscine

func LastRune(s string) rune {
	var x rune = 0
	for _, char := range s {
		x = char
	}
	return x
}
