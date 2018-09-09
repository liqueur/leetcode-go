package sol_Maximum_Subarray

import (
	"fmt"
	"math"
)

/*
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
 */


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func maxSubArray(nums []int) int {
	curSum, maxSum := 0, -math.MaxInt32
	for _, v := range nums {
		curSum = max(v, curSum + v)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}


func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
