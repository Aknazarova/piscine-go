package piscine

func IsPrintable(str string) bool {
	for _, char := range str {
		if !(char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z') {
			return false
		}
	}
	return true
}
