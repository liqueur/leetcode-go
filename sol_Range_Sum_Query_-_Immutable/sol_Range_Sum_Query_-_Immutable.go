package main

import "fmt"

type NumArray struct {
	content []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for x := i; x <= j; x++ {
		sum += this.content[x]
	}
	return sum
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	array := Constructor(nums)
	ret := array.SumRange(0, 5)
	fmt.Println(ret)
}
