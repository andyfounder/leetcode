package main

import (
	"fmt"
	"strings"
)

func main() {
	strArrays := []string{"10", "0001", "111001", "1", "0"}
	num0 := 5
	num1 := 3
	fmt.Println(findMaxForm(strArrays, num0, num1))
}

func findMaxForm(strs []string, m int, n int) int {
	array := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		array[i] = make([]int, n+1)
	}

	for _, str := range strs {
		count0 := strings.Count(str, "0")
		count1 := strings.Count(str, "1")
		if count0 > m || count1 > n {
			continue
		}
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				array[i][j] = max(array[i][j], 1+array[i-count0][j-count1])
			}
		}
	}
	return array[m][n]
}
func max(a1 int, a2 int) int {
	if a1 > a2 {
		a1, a2 = a2, a1
	}
	return a2
}
