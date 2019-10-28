package boolean

import "os"
import "github.com/01-edu/z01"

func main() {
	str := os.Args[1:]
	cnt := 0
	//	i := []rune (str)
	for i := range str {
		cnt = i + 1
	}
	if isEven(cnt) {
		printStr("I have an even number of arguments") //четный
	} else {
		printStr("I have an odd number of arguments") // нечетный
	}
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func printStr(str string) {
	arrayStr := []rune(str)
	counter := 0
	for i := range arrayStr {
		counter = i + 1
	}
	for i := 0; i < counter; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}
