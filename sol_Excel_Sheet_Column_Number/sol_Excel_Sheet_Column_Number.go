package main

import "fmt"

/*
Example 1:

Input: "A"
Output: 1
Example 2:

Input: "AB"
Output: 28
Example 3:

Input: "ZY"
Output: 701
 */


func titleToNumber(s string) int {
	ret := 0
	for _, v := range s {
		ret = int(v - 64) + ret * 26
	}

	return ret
}

func main() {
	ret := titleToNumber("AAA")
	fmt.Println(ret)
}
