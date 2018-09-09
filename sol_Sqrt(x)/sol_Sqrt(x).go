package sol_Sqrt_x_

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	pre := 1
	for i := 1; i < x; i++ {
		if i * i > x {
			return pre
		} else {
			pre = i
		}
	}

	return pre
}

func main() {
	fmt.Println(mySqrt(4))
}