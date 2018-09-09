package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n % 2 != 0 {
			return false
		}
		n /= 2
	}

	return true
}

func main() {
	fmt.Println(isPowerOfTwo(1))
}
