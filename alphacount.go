package piscine

func AlphaCount(str string) int {
	a := []rune(str)
	count := 0
	for range a {
		count++
	}
	return count
}
