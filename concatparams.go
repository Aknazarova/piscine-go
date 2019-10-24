package piscine

func ConcatParams(args []string) string {
	str := ""
	for index, char := range args {
		str += string(char)
		if index != len(args)-1 {
			str += "\n"
		}
	}
	return str
}
