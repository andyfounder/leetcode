package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var heights []int
	heights = []int{4, 2, 7, 6, 9, 14, 12}
	bricks := 5
	ladders := 1
	fmt.Println(furthestBuilding(heights, bricks, ladders))
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	usedLadder := &heightDiffPQ{}
	for i := 1; i < len(heights); i++ {
		needbricks := heights[i] - heights[i-1]
		if needbricks <= 0 {
			continue
		}
		if ladders > 0 {
			heap.Push(usedLadder, needbricks)
			ladders--
		} else {
			if len(*usedLadder) > 0 && needbricks > (*usedLadder)[0] {
				needbricks, (*usedLadder)[0] = (*usedLadder)[0], needbricks
				heap.Fix(usedLadder, 0)
			}
			bricks -= needbricks
			if bricks < 0 {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}

type heightDiffPQ []int

func (pq heightDiffPQ) Len() int {
	return len(pq)
}

func (pq heightDiffPQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq heightDiffPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *heightDiffPQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *heightDiffPQ) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}
