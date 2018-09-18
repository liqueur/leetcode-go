package main

import "fmt"

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func findErrorNums(nums []int) []int {
	ret := make([]int, 0)
	dup := -1
	mis := -1
	for _, v := range nums {
		if nums[abs(v)-1] < 0 {
			dup = abs(v)
		} else {
			nums[abs(v)-1] *= -1
		}
	}

	fmt.Println(nums)

	for i, v := range nums {
		if v > 0 {
			mis = i + 1
		}
	}

	ret = append(ret, dup, mis)

	return ret
}

/*
1 2 2 3 4 5 6 7 8 9]
0 1 0 0 0 0 0 0 0 0
*/
func main() {
	//ret := findErrorNums([]int{1,5,3,2,2,7,6,4,8,9})
	ret := findErrorNums([]int{2, 2})
	fmt.Println(ret)
}
