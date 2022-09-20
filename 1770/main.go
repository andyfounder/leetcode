package main

import (
	"math"
)

func main() {
	nums := []int{-5, -3, -3, -2, 7, 1}
	multipliers := []int{-10, -5, 3, 4, 6}
	maximumScore(nums, multipliers)
}

func maximumScore(nums []int, multipliers []int) int {
	dp := make([][]int, len(multipliers)+1)

	for i := 0; i < len(multipliers)+1; i++ {
		dp[i] = make([]int, len(multipliers)+1)
	}

	for i := 1; i < len(multipliers)+1; i++ {
		dp[i][0] = dp[i-1][0] + nums[i-1]*multipliers[i-1]
		dp[0][i] = dp[0][i-1] + nums[len(nums)-i]*multipliers[i-1]
	}

	for i := 1; i <= len(multipliers); i++ {
		for j := 1; i+j <= len(multipliers); j++ {
			dp[i][j] = Max(dp[i][j-1]+nums[len(nums)-j]*multipliers[i+j-1], dp[i-1][j]+nums[i-1]*multipliers[i+j-1])
		}
	}

	res := math.MinInt32

	for i := 0; i < len(multipliers)+1; i++ {
		res = Max(res, dp[i][len(multipliers)-i])
	}
	return res

}

func Max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// if len(multipliers) == 0 {
// 	return 0
// }
// l := nums[0]*multipliers[0] + maximumScore(nums[1:], multipliers[1:])
// r := nums[len(nums)-1]*multipliers[0] + maximumScore(nums[:len(nums)-1], multipliers[1:])
// if l > r {
// 	fmt.Println(l)
// 	return l
// } else {
// 	fmt.Println(r)
// 	return r
// }
