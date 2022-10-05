package main

func main() {
	colors := "abaac"
	neededTime := []int{1, 2, 3, 4, 5}
	minCost(colors, neededTime)
}

func minCost(colors string, neededTime []int) int {
	res, temp := 0, neededTime[0]
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			if temp < neededTime[i+1] {
				res += temp
				temp = neededTime[i+1]
			} else {
				res += neededTime[i+1]
			}
		} else {
			temp = neededTime[i+1]
		}
	}
	return res
}
