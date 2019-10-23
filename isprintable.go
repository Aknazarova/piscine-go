package piscine

func IsPrintable(str string) bool {
	for _, char := range str {
		if !(char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= 91 && char <= 96 || char >= '0' && char <= '9' ) {
			return false
		}
	}
	return true
}
