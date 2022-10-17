package main

func main() {
	string := "thequickbrownfoxjumpsoverthelazydog"
	checkIfPangram(string)
}

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	array := make([]bool, 26)
	for _, str := range sentence {
		array[str-'a'] = true
	}
	for _, arr := range array {
		if arr == false {
			return false
		}
	}
	return true
}
