package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world  "
	ss := reverseWords(s)
	fmt.Println(ss)
}

func reverseWords(s string) string {
	res := ""
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = strings.Trim(arr[i], " ")
		if arr[i] != "" {
			res += arr[i]
			res += " "
		}
	}
	res = res[:len(res)-1]
	return res
}
