package piscine

func SplitWhiteSpaces(str string) []string {
	a := make([]string, 1)
	for _, x := range str {
		if (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
			a[0] = str
		}
	}
	return a
}
