package main

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	transpose(matrix)
}

func transpose(matrix [][]int) [][]int {
	length1 := len(matrix)
	length2 := len(matrix[0])
	array := make([][]int, length2)
	for i := 0; i < length2; i++ {
		array[i] = make([]int, length1)
	}
	for i := 0; i < length1; i++ {
		for j := 0; j < length2; j++ {
			array[j][i] = matrix[i][j]
		}
	}
	return array
}
