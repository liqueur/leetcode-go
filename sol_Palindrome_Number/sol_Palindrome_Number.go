package sol_Palindrome_Number

import "fmt"

/*
Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 */

func isPalindrome(x int) bool {
	xx := x
	if x < 0 {
		xx = -x
	}

	res := 0
	for xx != 0 {
		res = res * 10 + xx % 10
		xx /= 10
	}

	if res == x {
		return true
	}

	return false
}


func main() {
	fmt.Println(isPalindrome(0))
}