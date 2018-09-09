package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	pre := []int{1}

	for i := 0; i < rowIndex; i++ {
		line := make([]int, 0)
		line = append(line, 1)
		for j := 0; j < len(pre) - 1; j++ {
			line = append(line, pre[j] + pre[j + 1])
		}
		line = append(line, 1)
		pre = line
	}

	return pre
}

func main() {
	fmt.Println(getRow(3))
}
