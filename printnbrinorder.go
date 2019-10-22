package piscine

func PrintNbrInOrder(n int) int {
	result := 0
	for n > 0 {
		remainder := n % 10
		result = result * 10
		result = result + remainder
		n = n / 10
	}
	return result
}
