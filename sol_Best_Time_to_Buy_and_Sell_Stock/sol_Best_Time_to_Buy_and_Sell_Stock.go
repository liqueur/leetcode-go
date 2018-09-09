package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minBuy := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		minBuy = min(prices[i], minBuy)
		profit = max(prices[i] - minBuy, profit)
	}

	return profit
}

/*
7,8,5,7,6,4
 */
func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
