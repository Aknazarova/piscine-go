package piscine

func IsUpper(str string) bool {
	for _, char := range str {
		if !(char >= 'A' && char <= 'Z') {
			return false
		}
	}
	return true
}
