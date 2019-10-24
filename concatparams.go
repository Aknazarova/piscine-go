package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, c := range args {
		str += string(c)
		if i != make([]string, args){
			str += "\n"
		}
	}
	return str
}
