package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}

	return true
}

func main() {
	ret := isPowerOfThree(9)
	fmt.Println(ret)
}
