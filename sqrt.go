package piscine

func Sqrt(nb int) int {
	result := 0
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			result = i
			break
		}
	}
	return result
}
