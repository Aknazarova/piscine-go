package piscine

func StrLen(str string) int {
	word := 0

	for _, a := range str  {
		if (a > 'A' && a < 'Z') || (a > 'a'&& a < 'z')
		word++
	}
	}
	return word
}
