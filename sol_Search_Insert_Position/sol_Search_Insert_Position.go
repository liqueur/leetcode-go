package sol_Search_Insert_Position

import "fmt"

/*
Input: [1,3,5,6], 5
Output: 2

Input: [1,3,5,6], 2
Output: 1


Input: [1,3,5,6], 7
Output: 4


Input: [1,3,5,6], 0
Output: 0
 */


func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
