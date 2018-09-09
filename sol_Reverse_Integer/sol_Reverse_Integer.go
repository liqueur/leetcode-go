package sol_Reverse_Integer

import (
	"fmt"
)

/*
Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
 */


func reverse(x int) int {
	res := 0
	for x != 0 {
		t := res * 10 + x % 10
		res = t
		x /= 10
	}

	if res >= -2147483648 && res <= 2147483647 {
		return res
	}

	return 0
}

func main() {
	fmt.Println(reverse(1534236469))
}
