package main

import (
	"fmt"
	"math"
)

func main() {
	s := "00110110"
	k := 2
	fmt.Println(hasAllCodes(s, k))
}

func hasAllCodes(s string, k int) bool {
	strLength, codeLen := len(s), k
	codeDict := make(map[string]bool)
	for start := 0; start < strLength-codeLen+1; start++ {
		cur_code := s[start : start+codeLen]
		if _, exist := codeDict[cur_code]; exist == false {
			codeDict[cur_code] = true
		}
	}
	return len(codeDict) == int(math.Pow(2.0, float64(k)))
}
