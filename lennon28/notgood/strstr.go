package main

import "fmt"

func Strstr1(haystack string, needle string, index, inpath, beforelen int, before string) int {
	fmt.Println(index)
	if len(needle) == 0 {
		if beforelen == inpath {
			return index
		} else {
			return -1
		}
	}
	if beforelen == 0 {
		beforelen = len(needle)
		if beforelen == 0 {
			return 0
		}
	}
	if len(haystack) == 0 {
		return -1
	}
	if needle[0] == haystack[0] {
		inpath += 1
		return Strstr1(haystack[1:], needle[1:], index, inpath, beforelen, before)
	} else {
		index += 1
		inpath = 0
		needle = before

		return Strstr1(haystack[1:], needle, index, inpath, beforelen, before)
	}
}

func strStr(haystack string, needle string) int {
	var index = 0
	var inpath = 0
	var beforelen = 0
	before := needle
	return Strstr1(haystack, needle, index, inpath, beforelen, before)
}

func main() {
	haystack := "mississippi"
	needle := "issip"
	re := strStr(haystack, needle)
	fmt.Println(re)
}
