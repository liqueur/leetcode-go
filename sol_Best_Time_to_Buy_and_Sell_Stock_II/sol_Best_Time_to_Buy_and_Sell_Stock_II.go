package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	for i := 0; i < len(prices) - 1; i++ {
		if prices[i] < prices[i + 1] {
			profit += prices[i + 1] - prices[i]
		}
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
}
