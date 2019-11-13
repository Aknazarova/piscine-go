package main

import "fmt"

func StrRev(s string) string {
	var ans string
	for _, c := range s {
		ans = string(c) + ans
	}
	return ans
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
