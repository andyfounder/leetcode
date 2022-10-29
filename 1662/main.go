func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var p1, p2 int
	var i, j int
	for p1 < len(word1) && p2 < len(word2) {
		if word1[p1][i] != word2[p2][j] {
			return false
		}
		i++
		if i == len(word1[p1]) {
			p1++
			i = 0
		}
		j++
		if j == len(word2[p2]) {
			p2++
			j = 0
		}
	}
	return len(word1) == p1 && len(word2) == p2
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var words1 string
	var words2 string
	for _, word := range word1 {
		words1 += word
	}
	for _, word := range word2 {
		words2 += word
	}
	if words1 == words2 {
		return true
	}
	return false
}