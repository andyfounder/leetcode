package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, -1, -2, 4, -7, 3}
	k := 2
	res := maxResult(nums, k)
	fmt.Println(res)
}
func maxResult(nums []int, k int) int {
	// dp := make([]int, len(nums))
	// if k > len(nums) {
	// 	k = len(nums)
	// }
	// dp[0] = nums[0]
	// for i := 1; i < len(dp); i++ {
	// 	dp[i] = math.MinInt32
	// }
	// for i := 1; i < len(nums); i++ {
	// 	left, tmp := max(0, i-k), math.MinInt32
	// 	for j := left; j < i; j++ {
	// 		tmp = max(tmp, dp[j])
	// 	}
	// 	dp[i] = nums[i] + tmp
	// }
	// return dp[len(nums)-1]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MinInt32
	}
	window := make([]int, k)
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + dp[window[0]]
		for len(window) > 0 && dp[window[len(window)-1]] <= dp[i] {
			window = window[:len(window)-1]
		}
		for len(window) > 0 && i-k >= window[0] {
			window = window[1:]
		}
		window = append(window, i)
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
