package main

import (
	"sort"
)

func main() {
	changed := []int{1, 3, 4, 2, 6, 7}
	findOriginalArray(changed)
}

// func findOriginalArray(changed []int) []int {
// 	if len(changed)%2 != 0 {
// 		return []int{}
// 	}
// 	sort.Slice(changed, func(i, j int) bool {
// 		return changed[i] > changed[j]
// 	})
// 	res := []int{}
// 	flag := false
// 	for len(changed) > 0 {
// 		flag = false
// 		for j := 0; j < len(changed); j++ {
// 			if changed[0] == changed[j]*2 {
// 				res = append(res, changed[j])
// 				temp := changed[j+1:]
// 				changed = changed[:j]
// 				changed = append(changed, temp...)
// 				changed = changed[1:]
// 				flag = true
// 				break
// 			}
// 		}
// 		if flag == false {
// 			return []int{}
// 		}
// 	}
// 	return res
// }

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return []int{}
	}
	sort.Ints(changed)
	m := make(map[int]int)
	res := []int{}
	for _, c := range changed {
		if m[c] == 0 {
			m[c*2]++
			res = append(res, c)
			continue
		}
		m[c]--
	}
	for _, mm := range m {
		if mm != 0 {
			return []int{}
		}
	}
	return res
}
