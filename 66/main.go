package main

import "fmt"

func main() {
	array := []int{9}
	res := plusOne(array)
	for _, r := range res {
		fmt.Println(r)
	}
}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	// newArray := make([]int, len(digits)+1)
	// newArray[0] = 1
	// return newArray
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
