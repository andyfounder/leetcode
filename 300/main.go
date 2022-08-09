package main

import "fmt"

func main() {
	array := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(array)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
