package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	fmt.Println(maxScore(array, k))
}

func maxScore(cardPoints []int, k int) int {
	// count := len(cardPoints) - k
	// total := 0
	// for _, c := range cardPoints {
	// 	total += c
	// }
	// left, right, sum, min := 0, count, 0, total
	// for right <= len(cardPoints) {
	// 	sum = 0
	// 	i := left
	// 	for i < right {
	// 		sum += cardPoints[i]
	// 		i++
	// 	}
	// 	if sum < min {
	// 		min = sum
	// 	}
	// 	left++
	// 	right++
	// }
	// return total - min
	windowSize, sum := len(cardPoints)-k, 0
	for _, c := range cardPoints[:windowSize] {
		sum += c
	}
	minSum := sum
	for i := windowSize; i < len(cardPoints); i++ {
		sum += cardPoints[i] - cardPoints[i-windowSize]
		if sum < minSum {
			minSum = sum
		}
	}
	total := 0
	for _, c := range cardPoints {
		total += c
	}
	return total - minSum
}
