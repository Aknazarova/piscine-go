package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, c := range args {
		str += string(c)
		if i != cap(args) - 1 {
			str += "\n"
		}
	}
	return str
}
