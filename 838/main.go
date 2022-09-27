package main

func main() {
	dominoes := "RR.L"
	pushDominoes(dominoes)
}

func pushDominoes(dominoes string) string {
	dominoes = "L" + dominoes + "R"
	res := ""
	for i, j := 0, 1; j < len(dominoes); j++ {
		if dominoes[j] == '.' {
			continue
		}
		if i > 0 {
			res += string(dominoes[i])
		}
		middle := j - i - 1
		if dominoes[i] == dominoes[j] {
			for a := 0; a < middle; a++ {
				res += string(dominoes[i])
			}
		} else if dominoes[i] == 'L' && dominoes[j] == 'R' {
			for a := 0; a < middle; a++ {
				res += string('.')
			}
		} else {
			for a := 0; a < middle/2; a++ {
				res += string('R')
			}
			for a := 0; a < middle%2; a++ {
				res += string('.')
			}
			for a := 0; a < middle/2; a++ {
				res += string('L')
			}
		}
		i = j
	}
	return res
}
