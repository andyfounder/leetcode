package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	fmt.Println(longestStrChain(array))
}

func longestStrChain(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	array := make([]int, 16+2)

	for i, w := range words {
		if array[len(w)] == 0 {
			array[len(w)] = i
		}
	}
	dp := make([]int, len(words))
	num := 0
	for i := len(words) - 1; i >= 0; i-- {
		dp[i] = 1
		for j := array[len(words[i])+1]; j < len(words) && len(words[j]) == len(words[i])+1; j++ {
			if check(words[i], words[j]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		num = max(num, dp[i])
	}
	return num
}

func check(a string, b string) bool {
	for i := 0; i < len(b); i++ {
		newString := b[:i] + b[i+1:]
		if a == newString {
			return true
		}
	}
	return false
}

func max(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	return a
}
