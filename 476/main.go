package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findComplement(5))
}
func findComplement(num int) int {
	n := int64(num)
	binary := strconv.FormatInt(n, 2)
	var reverse string
	for _, el := range binary {
		if el == '1' {
			reverse += "0"
		} else {
			reverse += "1"
		}
	}
	i, _ := strconv.ParseInt(reverse, 2, 64)
	return int(i)
}
