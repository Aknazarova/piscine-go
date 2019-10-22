package piscine

func ToLower(s string) string {
	var a string
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			a += string(c + 32)
		} else {
			a += string(c)
		}
	}
	return a
}
