package main

import "fmt"

func firstUniqChar(s string) int {
	smap := make(map[interface{}]int, 0)
	for _, v := range s {
		smap[v] += 1
	}

	for i, v := range s {
		if smap[v] <= 1 {
			return i
		}
	}

	return -1
}

func main() {
	ret := firstUniqChar("leetcode")
	fmt.Println(ret)
}
