package main

import "sort"

func main() {
	coins := []int{1, 4, 10, 3, 1}
	getMaximumConsecutive(coins)
}

func getMaximumConsecutive(coins []int) int {
	// sort.Ints(coins)
	// left, right, max := 0, 0, 1
	// for _, coin := range coins {
	// 	if left+coin <= right+1 {
	// 		right += coin
	// 	} else {
	// 		left, right = coin, coin
	// 	}
	// 	if right-left+1 > max {
	// 		max = right - left + 1
	// 	}
	// }
	// return max

	sort.Ints(coins)
	temp := 0
	for _, coin := range coins {
		if temp+1 < coin {
			break
		}
		temp += coin
	}
	return temp + 1
}
