package main

import (
	"fmt"
)

func main() {
	arrays := []int{2, 2, 1}
	fmt.Println(singleNumber(arrays))
}
func singleNumber(nums []int) int {
	temp := 0
	for _, n := range nums {
		temp ^= n
	}
	return temp
}
