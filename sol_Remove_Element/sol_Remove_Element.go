package sol_Remove_Element

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	n := 0

	for i, v := range nums {
		if v == val {
			for j := i; j < len(nums); j++ {
				if nums[j] != nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					n = i + 1
					break
				}
			}
		} else {
			n += 1
		}
	}
	return n
}

func main() {
	//nums, val := []int{3, 2, 2, 3}, 3
	//nums, val := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	//nums, val := []int{1}, 2
	nums, val := []int{1}, 1
	n := removeElement(nums, val)
	fmt.Println(nums, n)
}
