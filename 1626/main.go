package main

import (
	"sort"
)

func main() {
	scores := []int{4, 5, 6, 5}
	ages := []int{2, 1, 2, 1}
	bestTeamScore(scores, ages)
}

func bestTeamScore(scores []int, ages []int) int {
	res := 0
	array := [][2]int{}
	for i := 0; i < len(scores); i++ {
		array = append(array, [2]int{})
		array[i][0] = scores[i]
		array[i][1] = ages[i]
	}
	sort.Slice(array, func(i, j int) bool {
		if array[i][1] < array[j][1] {
			return true
		}
		if array[i][1] == array[j][1] && array[i][0] < array[j][0] {
			return true
		}
		return false
	})
	dp := []int{}
	for i := 0; i < len(scores); i++ {
		dp = append(dp, 0)
	}
	for i := 0; i < len(array); i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j][0] <= array[i][0] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += array[i][0]
		res = max(res, dp[i])
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
