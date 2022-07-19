package main

import "fmt"

func main() {
	str := "QWEZX"
	fmt.Println(titleToNumber(str))
}

func titleToNumber(columnTitle string) int {
	charArrays := []rune(columnTitle)
	temp := 1
	sum := 0
	for i := len(charArrays) - 1; i >= 0; i-- {
		value := (int(charArrays[i]-'A') + 1) * temp
		sum += value
		temp *= 26
	}
	return sum
}
