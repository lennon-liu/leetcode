package main

import (
	"fmt"
	"sort"
)
import "math"

func twodivde(nums []int, target int) int {
	//fmt.Println(nums)
	if len(nums) == 2 {
		if math.Abs(float64(nums[0]-target))-math.Abs(float64(nums[1]-target)) > 0 {
			return nums[1]
		} else {
			return nums[0]
		}
	}
	length := len(nums) - 1
	if length%2 == 0 {
		length = length / 2
	} else {
		length = (length - 1) / 2
	}
	if target == nums[length] {
		return target
	} else if target > nums[length] {
		return twodivde(nums[length:], target)
	} else {
		return twodivde(nums[:length+1], target)
	}
}

func threeSumClosest(nums []int, target int) int {
	result := 0
	lastx := float64(10000)

	ttt := 0
	sort.Ints(nums)
	for i1 := 0; i1 <= len(nums)-3; i1++ {
		lastnum3 := 9999
		for i2 := i1 + 1; i2 <= len(nums)-2; i2++ {
			num3 := (target - nums[i1]) - nums[i2]
			if num3 > nums[len(nums)-1] {
				ttt = nums[len(nums)-1] + nums[i2] + nums[i1]
			} else if num3 < nums[i2+1] {
				ttt = nums[i2+1] + nums[i2] + nums[i1]
			} else {
				ttt = twodivde(nums, num3) + nums[i2] + nums[i1]
			}
			if ttt-target == 0 {
				return ttt
			}
			if math.Abs(float64(ttt-target))-math.Abs(lastx) < 0 {
				lastx = float64(ttt - target)
				result = ttt
			}
			if ttt-target > 0 && ttt-lastnum3 > 0 {
				break
			} else {
				lastnum3 = ttt
			}
		}
	}
	return result
}

func main() {
	testlist := []int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33}
	fmt.Println(threeSumClosest(testlist, 0))
}
