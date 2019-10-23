package printprogramname

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	z01.PrintRune(os.Args[0])

}
