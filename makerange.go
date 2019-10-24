package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var a []int = nil
		return a
	} else {
		a := make([]int, max-min) //выводит обнуленный массив
		j := 0
		for i := min; i < max; i++ {
			a[j] = i
			j++
		}
		return a
	}
}
