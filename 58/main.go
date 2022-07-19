package main

import (
	"fmt"
)

func main() {
	str := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(str))
}
func lengthOfLastWord(s string) int {
	// return len(strings.TrimSpace(s)) - strings.LastIndex(strings.TrimSpace(s), " ") - 1
	suf := len(s) - 1
	for suf >= 0 && s[suf] == ' ' {
		suf--
	}
	if suf < 0 {
		suf = 0
	}
	pre := suf
	for pre >= 0 && s[pre] != ' ' {
		pre--
	}
	return suf - pre
}
