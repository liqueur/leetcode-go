package sol_Climbing_Stairs

import "fmt"

/*
Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 */

func climbStairs(n int) int {
	s1, s2 := 1, 2

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	for i := 3; i <= n; i++ {
		s3 := s1 + s2
		s1, s2 = s2, s3
	}

	return s2
}

func main() {
	fmt.Println(climbStairs(5))
}
