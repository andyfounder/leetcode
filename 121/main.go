package main

import "fmt"

func main() {
	arrays := []int{2, 1, 3, 6}
	fmt.Println(maxProfit(arrays))
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	profit := 0
	pre := 0
	suf := 0
	for i := 0; i < len(prices)-1; i++ {
		suf = i + 1
		if profit < prices[suf]-prices[pre] {
			profit = prices[suf] - prices[pre]
		}
		if prices[i+1] < prices[pre] {
			pre = i + 1
			suf = pre
		}
	}
	return profit
}
