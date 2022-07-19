package main

import (
	"fmt"
)

func main() {
	num := 152146
	fmt.Println(convertToTitle(num))
}

func convertToTitle(columnNumber int) string {
	var str string
	for columnNumber > 0 {
		value := (columnNumber - 1) % 26
		columnNumber = (columnNumber - 1) / 26
		tempstr := string(rune(value + 'A'))
		tempstr += str
		str = tempstr
	}
	return str
}
