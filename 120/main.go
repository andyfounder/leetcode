package main

import (
	"fmt"
	"math"
)

func main() {
	array := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(array))
}

// func minimumTotal(triangle [][]int) int {
// 	for i := len(triangle) - 2; i >= 0; i-- {
// 		for j := 0; j < len(triangle[i]); j++ {
// 			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
// 		}
// 	}
// 	return triangle[0][0]
// }

// func min(a int, b int) int {
// 	if a > b {
// 		a, b = b, a
// 	}
// 	return a
// }

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp, minNum, index := make([]int, len(triangle[len(triangle)-1])), math.MaxInt64, 0
	for ; index < len(triangle[0]); index++ {
		dp[index] = triangle[0][index]
	}
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 {
				dp[j] += dp[j-1] + triangle[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		if minNum > dp[i] {
			minNum = dp[i]
		}
	}
	return minNum
}

func min(a int, b int) int {
	if a > b {
		a, b = b, a
	}
	return a
}
