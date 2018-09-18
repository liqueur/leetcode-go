package main

func intersection(nums1 []int, nums2 []int) []int {
	a := make(map[int]int, 0)
	b := make(map[int]int, 0)
	ret := make([]int, 0)

	for _, v := range nums1 {
		a[v] += 1
	}

	for _, v := range nums2 {
		_, in := a[v]
		if in {
			b[v] += 1
		}
	}

	for k, _ := range b {
		ret = append(ret, k)
	}

	return ret
}

func main() {

}
