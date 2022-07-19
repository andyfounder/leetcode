package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [][]int{{1, 3}, {5, 5}, {2, 5}, {4, 2}, {4, 1}, {3, 1}, {2, 2}, {1, 3}, {2, 5}, {3, 2}}
	max := 35
	fmt.Println(maximumUnits(array, max))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	// min := 0
	// for m := 0; m < len(boxTypes); m++ {
	// 	if boxTypes[min][1] > boxTypes[m][1] {
	// 		min = m
	// 	}
	// }
	// i, res := 0, 0
	// for i < truckSize {
	// 	temp := boxTypes[min][1]
	// 	j := 0
	// 	max := min
	// 	for j = 0; j < len(boxTypes); j++ {
	// 		if temp <= boxTypes[j][1] && boxTypes[j][0] > 0 {
	// 			temp = boxTypes[j][1]
	// 			max = j
	// 		}
	// 	}
	// 	if max == min && boxTypes[max][0] == 0 {
	// 		break
	// 	}
	// 	res += boxTypes[max][1]
	// 	i++
	// 	boxTypes[max][0]--
	// }
	// return res
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	for i := 0; truckSize > 0 && i < len(boxTypes); i++ {
		if truckSize >= boxTypes[i][0] {
			res += boxTypes[i][1] * boxTypes[i][0]
			truckSize -= boxTypes[i][0]
			boxTypes[i][0] = 0
		} else {
			res += boxTypes[i][1] * truckSize
			truckSize = 0
			// boxTypes[i][0] -= truckSize
		}
	}
	return res
}
