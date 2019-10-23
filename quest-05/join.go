package piscine

func Join(strs []string, sep string) string {
	var a string
	var count int = 0
	for range strs {
		count++
	}
	for b := range strs {
		a += strs[b]
		if b != count-1 { // if it is last word
			a += sep // print ':'
		}
	}
	return a
}
