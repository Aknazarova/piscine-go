package piscine

func IsPrintable(str string) bool {
	for _, char := range str {
		if char >= 7 && char <= 13 {
			return false
		}
	}
	return true
}
