package main

import "fmt"

func main() {
	m := 3
	n := 7
	// array := [2][3]int{{1, 1, 1}, {1, 1, 1}}
	// fmt.Println(len(array))
	res := uniquePaths(m, n)
	fmt.Println(res)
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//recursion
//dynamic programming
//make
//[][]int len
