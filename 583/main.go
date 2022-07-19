package main

import (
	"fmt"
)

func main() {
	s1 := "leetcode"
	s2 := "etco"
	fmt.Println(minDistance(s1, s2))
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
func min(a int, b int) (c int) {
	if a < b {
		c = a
	} else {
		c = b
	}
	return
}
