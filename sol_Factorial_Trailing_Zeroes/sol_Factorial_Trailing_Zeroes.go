package main

import "fmt"

func trailingZeroes(n int) int {
	if n / 5 == 0 {
		return 0
	} else {
		return n / 5 + trailingZeroes(n / 5)
	}
}


func main() {
	fmt.Println(trailingZeroes(225))
}
