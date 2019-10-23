package printprogramname

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argument := os.Args
	z01.PrintRune(argument)
}
