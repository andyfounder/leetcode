package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 4, 5, 10}
	res := numFactoredBinaryTrees(arr)
	fmt.Println(res)
}

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)

}
