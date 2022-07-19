package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 2, 4}
	sortArrayByParity(nums)
	for _, n := range nums {
		fmt.Println(n)
	}
}
func sortArrayByParity(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
