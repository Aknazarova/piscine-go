package piscine

func main() {

	s := "0000000012345"

	/** converting the s variable into an int using Atoi method */
	n, err := BasicAtoi(s)
	if err == nil {
		fmt.PrintRune(n)
	}

}
