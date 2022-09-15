package main

func main() {
	n := 4
	edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	distanceThreshold := 4
	findTheCity(n, edges, distanceThreshold)
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			distance[i][j] = 10001
		}
	}
	for _, edge := range edges {
		distance[edge[0]][edge[1]] = edge[2]
		distance[edge[1]][edge[0]] = edge[2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distance[i][k]+distance[k][j] < distance[i][j] && i != j {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}
	res, min, count := 0, 101, 0
	for i := 0; i < n; i++ {
		count = 0
		for j := 0; j < n; j++ {
			if distance[i][j] <= distanceThreshold {
				count++
			}
		}
		if count <= min {
			min = count
			res = i
		}
	}
	return res
}
