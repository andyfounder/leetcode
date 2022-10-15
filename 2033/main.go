func minOperations(grid [][]int, x int) int {
	num := len(grid) * len(grid[0])
	array := make([]int, 0)
	for _, gri := range grid {
		for _, gr := range gri {
			array = append(array, gr)
			if (gr-grid[0][0])%x != 0 {
				return -1
			}
		}
	}
	sort.Ints(array)
	result := 0
	for _, a := range array {
		result += abs(a, array[num/2]) / x
	}
	return result
}

func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}