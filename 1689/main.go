package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "32"
	fmt.Println(minPartitions(s))
}

func minPartitions(n string) int {
	slice := strings.Split(n, "")
	res := 0
	for _, s := range slice {
		i, _ := strconv.Atoi(s)
		if i > res {
			res = i
		}
	}
	return res
}
