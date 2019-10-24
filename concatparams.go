package piscine

func ConcatParams(args []string) string {
	var str string
	var count int = 0

	for range args {
		count++
	}
	for c := range args {
		str += args[c]
		if c != count-1 {
			str += "\n"
		}
	}
	return str
}
