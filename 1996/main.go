package main

import "sort"

func main() {
	array := [][]int{{5, 5}, {6, 3}, {3, 6}}
	numberOfWeakCharacters(array)
}

func numberOfWeakCharacters(properties [][]int) int {
	if len(properties) < 2 {
		return 0
	}
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] > properties[j][0] {
			return true
		}
		if properties[i][0] == properties[j][0] && properties[i][1] < properties[j][1] {
			return true
		}
		return false
	})
	count := 0
	temp := 0
	for i := 0; i < len(properties); i++ {
		if temp > properties[i][1] {
			count++
		}
		if temp < properties[i][1] {
			temp = properties[i][1]
		}
	}
	return count
}
