package main

func main() {
	strArray := []string{"b==a", "a==b", "b==c", "c==d"}
	equationsPossible(strArray)
}

func equationsPossible(equations []string) bool {
	array := make([]int, 26)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '=' {
			index1 := int(equations[i][0] - 'a')
			index2 := int(equations[i][3] - 'a')
			union(array, index1, index2)
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			if find(array, index1) == find(array, index2) {
				return false
			}
		}
	}
	return true
}

func union(array []int, index1, index2 int) {
	array[find(array, index1)] = find(array, index2)
}

func find(array []int, index int) int {
	for array[index] != index {
		array[index] = array[array[index]]
		index = array[index]
	}
	return index
}
