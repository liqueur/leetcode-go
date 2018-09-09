package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := []int{nums[0], max(nums[0], nums[1])}

	for i := 2; i < len(nums); i++ {
		dp = append(dp, max(nums[i] + dp[i - 2], dp[i - 1]))
	}

	return dp[len(nums) - 1]
}


// 1, 0, 0, 100, 0
func main() {
	fmt.Println(rob([]int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}))
}
