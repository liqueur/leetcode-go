package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	dup := make(map[int]int, 0)
	for i, v := range nums {
		di, in := dup[v]
		if in && i - di <= k{
			return true
		}
		dup[v] = i
	}

	return false
}

func main() {
	ret := containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2)
	fmt.Println(ret)
}
