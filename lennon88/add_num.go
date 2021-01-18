package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var result = make([]int, m)
	i, j := 0, 0
	for {
		if nums1[i] > nums2[j] {
			result = append(result, nums2[j])
			i += 1
		} else if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			result = append(result, nums2[j])
			i += 1
			j += 1
		} else {
			result = append(result, nums2[j])
			j += 1
		}
		if i+j == m {
			break
		}
		fmt.Println(result)
	}

}

func main() {
	nums1, nums2 := []int{1, 2, 0}, []int{2}
	m, n := 3, 2
	merge(nums1, m, nums2, n)
}
