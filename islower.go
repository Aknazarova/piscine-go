package piscine

func IsLower(str string) bool {
	for _, char := range str {
		if !(char >= 'a' && char <= 'z') {
			return false
		}
	}
	return true
}
