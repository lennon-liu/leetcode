package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var result [110 + 110]int
	length := (len(num1) - 1) + (len(num2) - 1) + 1
	for i1 := len(num1) - 1; i1 >= 0; i1-- {
		for i2 := len(num2) - 1; i2 >= 0; i2-- {
			numx := (int(num1[i1]) - 48) * (int(num2[i2]) - 48)
			if numx >= 10 {
				result[i1+i2+1] += numx % 10
				result[i1+i2] += numx / 10
			} else {
				result[i1+i2+1] += numx
			}
			for i3 := i1 + i2 + 1; i3 >= 0; i3-- {
				temp := result[i3]
				if temp >= 10 {
					result[i3] = temp % 10
					result[i3-1] += temp / 10
				} else {
					break
				}

			}
		}
	}
	str := ""
	for i3 := 0; i3 <= length; i3++ {
		str += string(result[i3] + 48)
	}
	if str[0] == 48 {
		str = str[1:]
	}
	return str
}

func main() {
	num1 := "999"
	num2 := "999"
	fmt.Println(multiply(num1, num2))
	fmt.Println(999 * 999)
}
