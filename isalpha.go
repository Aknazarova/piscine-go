package piscine

func IsAlpha(str string) bool {
	for _, char := range str {
		if! (char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}
