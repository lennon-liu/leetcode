package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	max := len(digits) - 1
	for {
		if digits[max] == 10 {
			if max == 0 {
				digits[max] = 0
				digits = append([]int{1}, digits...)
			} else {
				digits[max] = 0
				digits[max-1] = digits[max-1] + 1
			}
		} else {
			break
		}
		if max < 1 {
			break
		}
		max -= 1
	}
	return digits
}

func main() {
	digits := []int{
		0, 0, 1, 0, 0, 0,
	}
	result := plusOne(digits)
	fmt.Println(result)
}

//[3,3,9,0,7,4,1,6,1,9,1,5,4,4,7,6,6,6,3]
//[7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,7]
