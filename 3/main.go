package main

import (
	"fmt"
)

func main() {
	str := "abc"
	l := lengthOfLongestSubstring(str)
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {

	// O(n^2)
	// strArray := strings.Split(s, "")
	// pre := 0
	// now := 0
	// max := 0
	// for i := 0; i < len(strArray); i++ {
	// 	now = i
	// 	for j := pre; j < now; j++ {
	// 		if strArray[j] == strArray[i] {
	// 			pre = j + 1
	// 			break
	// 		}
	// 	}
	// 	if now-pre+1 > max {
	// 		max = now - pre + 1
	// 	}
	// }
	// return max

	// O(n)
	// if len(s) == 0 {
	// 	return 0
	// }
	// bitArray := [256]bool{}
	// result, left, right := 0, 0, 0
	// for left < len(s) {
	// 	if bitArray[s[right]] {
	// 		bitArray[s[left]] = false
	// 		left++
	// 	} else {
	// 		bitArray[s[right]] = true
	// 		right++
	// 	}
	// 	if right-left > result {
	// 		result = right - left
	// 	}
	// 	if right == len(s) {
	// 		break
	// 	}
	// }
	// return result

	//O(n)
	// if len(s) == 0 {
	// 	return 0
	// }

	// array := [127]int{}
	// res, left, right := 0, 0, -1
	// for left < len(s) {
	// 	if array[s[right+1]] == 0 {
	// 		array[s[right+1]]++
	// 		right++
	// 	} else {
	// 		array[s[left]]--
	// 		left++
	// 	}
	// 	if right-left+1 > res {
	// 		res = right - left + 1
	// 	}
	// 	if right == len(s)-1 {
	// 		break
	// 	}
	// }
	// return res

	res, left, right := 0, 0, 0
	hashMap := make(map[byte]int, len(s))
	for left < len(s) {
		if h, ok := hashMap[s[left]]; ok && h >= right {
			right = h + 1
		}
		hashMap[s[left]] = left
		left++
		if left-right > res {
			res = left - right
		}
	}
	return res

	// ss := strings.Split(s, "")
	// length := len(ss)
	// max := 0
	// pre := 0
	// now := 0
	// for i := 0; i < length; i++ {
	// 	now = i
	// 	for j := pre; j < now; j++ {
	// 		if ss[j] == ss[i] {
	// 			pre = j + 1
	// 			continue
	// 		}
	// 	}
	// 	if now-pre+1 > max {
	// 		max = now - pre + 1
	// 	}
	// }
	// return max
}
