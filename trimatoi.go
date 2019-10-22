package piscine

func TrimAtoi(s string) int {
	sign := true
	arg := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			arg *= 10
			arg += int(char) - 48 // 48 тут равно 0,int(char) поменяет на ascii code
		} else if char == '-' && arg == 0 {
			sign = false
		}
	}
	if sign {
		return arg
	} else {
		return (-1) * arg
	}
}
