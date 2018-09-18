package main

import "fmt"

func findNthDigit(n int) int {
	len := 1
	cnt := 9
	start := 1
	for n > len*cnt {
		n -= len * cnt
		len++
		cnt *= 10
		start *= 10
	}

	start += (n - 1) / len
	t := fmt.Sprintf("%v", start)
	return int(t[(n-1)%len] - '0')
}

func main() {
	ret := findNthDigit(3)
	fmt.Println(ret)
}
