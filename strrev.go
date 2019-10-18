package strrev

func StrRev(s string) string {
	var count int = -1
	for range s {
		count++
	}
	sX := []byte(s)
	for i := 0; i <= count; i++ {
		sX[i] = s[count-i]
	}
	return string(sX)
}
