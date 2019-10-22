package piscine

func Compare(a, b string) int {
	str := []rune(a)
	c := []rune(b)
	n := 0
	m := 0
	for range str {
		n++
	}
	for range c {
		m++
	}
	for i := 0; i <= n-m; i++ {
		if b == a {
			return 0
		} else if a > b {
			return 1
		}
	}
	return -1
}
