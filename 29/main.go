package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 3
	fmt.Println(divide(a, b))
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend > math.MaxInt32 {
		dividend = math.MaxInt32
	}
	sign := -1
	left := 0
	right := dividend - 1
	mid := 0
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}
	for left <= right {
		mid = left + (right-left)>>1
		if (mid+1)*abs(divisor) < abs(dividend) {
			left = mid + 1
		} else if mid*abs(divisor) > abs(dividend) {
			right = mid - 1
		} else {
			if mid > math.MaxInt32 {
				return sign * math.MaxInt32
			}
			if mid < math.MinInt32 {
				return sign * math.MinInt32
			}
			return sign * mid
		}
	}
	return 0
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
