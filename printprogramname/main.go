package printprogramname

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	num := os.Args
	for _, c := range num[0] {
		z01.PrintRune(c)
	}
	z01.PrintRune(10)
}
