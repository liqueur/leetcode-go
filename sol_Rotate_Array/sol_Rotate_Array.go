package main

import "fmt"

func rotate(nums []int, k int)  {
	for i := 0; i < k; i++ {
		v := nums[len(nums) - 1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j - 1]
		}
		nums[0] = v
	}
}

func main() {
	nums := []int{-1,-100,3,99}
	rotate(nums, 3)
	fmt.Println(nums)
}
