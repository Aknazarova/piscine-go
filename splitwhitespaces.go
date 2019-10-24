package piscine
func SplitWhiteSpaces(str string) []string {
	i := 0
	sign := true
	for _, char := range str{
		if !(char == ' ' || char == '\t' || char == '\n') {
				sign = true
		} else { 
			if sign {
				i++
				sign = false
			}
		}
	}
	a := make ([]string, i+1)
	j := 0
	mark := true
	for j != i {
		for _, char := range str {
			if !(char == ' ' || char == '\t' || char == '\n') {
				a[j] += string (char)
				mark = true
			} else {
				if mark {
					j++
					mark = false 
				}
			}
		}
	}
	return a
}