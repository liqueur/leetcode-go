package main

func containsDuplicate(nums []int) bool {
	dup := make(map[int]int, 0)
	for _, v := range nums {
		_, in := dup[v]
		if in {
			return true
		} else {
			dup[v] = 1
		}
	}
	return false
}

func main() {
	
}
