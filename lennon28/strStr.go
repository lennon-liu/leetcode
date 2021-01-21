package main

import "fmt"

func strStr(haystack string, needle string) int {
	i := 0
	for {
		if i > len(haystack) || len(haystack[i:]) < len(needle) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
		i += 1
	}
}

func main() {
	haystack := ""
	needle := ""
	re := strStr(haystack, needle)
	fmt.Println(re)
}
