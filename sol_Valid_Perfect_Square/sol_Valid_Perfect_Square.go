package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	a := num / 2
	for a*a > num {
		a = a / 2
	}

	for b := a; b <= 2*a; b++ {
		if b*b == num {
			return true
		}
	}

	return false
}

func main() {
	ret := isPerfectSquare(14)
	fmt.Println(ret)
}
