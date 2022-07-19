package main

import (
	"fmt"
)

func main() {
	words := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(words))
}

func minimumLengthEncoding(words []string) int {
	count := 0
	m := make(map[string]int)
	for _, word := range words {
		m[word] = len(word)
	}
	for w := range m {
		for i := 1; i < len(w); i++ {
			delete(m, w[i:])
		}
	}
	for w := range m {
		count += len(w) + 1
	}
	return count
}
