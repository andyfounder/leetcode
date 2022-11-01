package main

func main() {
	grid := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	findBall(grid)
}

func findBall(grid [][]int) []int {
	n := len(grid[0])
	result := make([]int, n)
	for i := range result {
		column := i
		for _, row := range grid {
			dir := row[column]
			column += dir
			if column < 0 || column == n || row[column] != dir {
				column = -1
				break
			}
		}
		result[i] = column
	}
	return result
}
