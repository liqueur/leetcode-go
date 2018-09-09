package sol_Remove_Duplicates_from_Sorted_Array


import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	n := 1
	for i, v := range nums {
		if i > 0 && v != nums[n - 1] {
			nums[n] = v
			n += 1
		}
	}

	return n
}

func main() {
	fmt.Println(removeDuplicates([]int{1,1,2}))
}
