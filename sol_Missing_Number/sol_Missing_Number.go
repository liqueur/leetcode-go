package main

import "fmt"

func missingNumber(nums []int) int {
	sum := 0
	n := len(nums)

	for _, v := range nums {
		sum += v
	}

	return int(n*(n+1)/2 - sum)
}

func main() {
	ret := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	fmt.Println(ret)
}
