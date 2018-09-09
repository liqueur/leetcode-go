package main

import "fmt"

/*
Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
 */

func isHappy(n int) bool {
	previous := make(map[int]int, 0)
	for {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if sum == 1 {
			return true
		}
		_, in := previous[sum]
		if in {
			return false
		} else {
			previous[sum] = 1
		}
		n = sum
	}
}

func main() {
	fmt.Println(isHappy(0))
}
