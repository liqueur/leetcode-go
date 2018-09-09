package main

import "fmt"

func generate(numRows int) [][]int {
	ret := make([][]int, 0)

	if numRows < 0 {
		return ret
	}

	ret = append(ret, []int{1})

	for i := 0; i < numRows - 1; i++ {
		pre := ret[len(ret) - 1]
		line := make([]int, 0)
		line = append(line, 1)
		for j := 0; j < len(pre) - 1; j++ {
			line = append(line, pre[j] + pre[j + 1])
		}
		line = append(line, 1)
		ret = append(ret, line)
	}

	return ret
}

func main() {
	fmt.Println(generate(3))
}
