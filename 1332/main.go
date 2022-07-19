package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("a"))
}

func removePalindromeSub(s string) int {
	// for i := 0; i < len(s)/2; i++ {
	// 	if s[i] != s[len(s)-1-i] {
	// 		return 2
	// 	}
	// }
	// return 1
	array := make(map[rune]int)
	for _, char := range s {
		if _, exist := array[char]; exist {
			array[char]++
		} else {
			array[char] = 1
		}
	}
	if isPalindrome(s) {
		return 1
	}
	return len(array)
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[right] != s[left] {
			return false
		}
		left++
		right--
	}
	return true
}
