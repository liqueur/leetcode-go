package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func intersect(nums1 []int, nums2 []int) []int {
	a := make(map[int]int, 0)
	b := make(map[int]int, 0)
	ret := make([]int, 0)

	for _, v := range nums1 {
		a[v] += 1
	}

	for _, v := range nums2 {
		_, in := a[v]
		if in {
			b[v] = min(b[v]+1, a[v])
		}
	}

	for k, v := range b {
		for i := 0; i < v; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

func main() {
	ret := intersect([]int{4, 9, 5, 9}, []int{9, 4, 9, 8, 4})
	fmt.Println(ret)
}
