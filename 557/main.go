package main

import (
	"strings"
)

func main() {
	str := "Let's take LeetCode contest"
	reverseWords(str)
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = reverse(s)
	}
	return strings.Join(ss, " ")

	// res, i := "", strings.Index(s, " ")
	// for i != -1 {
	// 	for j := i - 1; j >= 0; j-- {
	// 		res += s[j : j+1]
	// 	}
	// 	res += " "
	// 	s = s[i+1:]
	// 	i = strings.Index(s, " ")
	// }
	// res += s[len(s)-1:]
	// for i := len(s) - 2; i >= 0; i-- {
	// 	res += s[i : i+1]
	// }
	// return res
}

func reverse(s string) string {
	bytes, i, j := []byte(s), 0, len(s)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
