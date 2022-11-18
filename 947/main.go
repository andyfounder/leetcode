package main

import "fmt"

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	res := removeStones(stones)
	fmt.Println(res)
}

func removeStones(stones [][]int) int {
	fa, rank := map[int]int{}, map[int]int{}
	for _, stone := range stones {
		union(stone[0], stone[1]+10001, rank, fa)
	}
	res := len(stones)
	for x, fx := range fa {
		if x == fx {
			res--
		}
	}
	return res
}
func union(x int, y int, rank map[int]int, fa map[int]int) {
	fx, fy := find(x, rank, fa), find(y, rank, fa)
	if fx == fy {
		return
	}
	if rank[fx] < rank[fy] {
		fx, fy = fy, fx
	}
	rank[fx] += rank[fy]
	fa[fy] = fx
}

func find(x int, rank map[int]int, fa map[int]int) int {
	if _, has := fa[x]; !has {
		fa[x], rank[x] = x, 1
	}
	if fa[x] != x {
		fa[x] = find(fa[x], rank, fa)
	}
	return fa[x]
}
