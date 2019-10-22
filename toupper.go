package piscine

func ToUpper(s string) string {
	var a string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			a += string(c - 32)
		} else {
			a += string(c)
		}
	}
	return a
}
