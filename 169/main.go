package main

import "fmt"

func main() {
	array := []int{3, 2, 3}
	fmt.Println(majorityElement(array))
}

func majorityElement(nums []int) int {
	// temp := 0
	// value := nums[0]
	// for _, num := range nums {
	// 	if num == value {
	// 		temp++
	// 	} else {
	// 		temp--
	// 		if temp == 0 {
	// 			value = num
	// 			temp = 1
	// 		}
	// 	}
	// }
	// return value

	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}
