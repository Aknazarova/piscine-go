package piscine

func AlphaCount(str string) int {
	count := 0
	for _, x := range str {
		if (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
			count++
		}
	}
	return count
}
