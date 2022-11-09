package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "abc"
	for _, s := range str {
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Println(reflect.TypeOf(str[i]))
	}
	s := "aA"
	s = makeGood(s)
	fmt.Println(s)
}

func makeGood(s string) string {
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1]+32 || s[i] == s[i+1]-32 {
				s = s[:i] + s[i+2:]
				flag = true
			}
		}
	}
	return s
}
