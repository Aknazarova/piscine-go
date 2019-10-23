package piscine

func Join(strs []string, sep string) string {
	var a string
	for b := range strs {
		a += strs[b] + sep
	}
	return a
}

