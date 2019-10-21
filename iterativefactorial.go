package main

import "fmt"

func IterativeFactorial(nb int) int {

	if nb >= 0 {
		result := 1
		for i := 2; i <= nb; i++{
		result = result * i
		}
		return result
	} else {
	return 0
	}
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
