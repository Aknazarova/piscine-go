package piscine

func str (s string) {
	for _, char := rang s {
		z01.PrintRune(char)
	}
}
func PrintWordsTables(table []string) {
	counter := 0
	for range table {
		counter++
	}	
	for _, strstr := range table {
		str (strstr)
		z01.PrintRune('\n')
	}
}
