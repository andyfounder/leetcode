package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "bbcebab"
	fmt.Println(minDeletions(s))
}

func minDeletions(s string) int {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	i, temp, array := 0, slice[0], make([]int, 1)
	for _, ss := range slice {
		if ss != temp {
			array = append(array, 1)
			temp = ss
			i++
		} else {
			array[i]++
		}
	}
	sort.Ints(array)
	if len(array) == 1 {
		return 0
	}
	now := array[len(array)-1]
	res := 0
	for j := len(array) - 2; j >= 0; j-- {
		if array[j] < now {
			now = array[j]
			continue
		} else {
			if now == 0 {
				res += array[j]
				array[j] = 0
				continue
			}
			for array[j] >= now {
				array[j]--
				res++
			}
			now = array[j]
		}
	}
	return res
}
