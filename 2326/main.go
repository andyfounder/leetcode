package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	m := 3
	n := 5
	node0 := new(ListNode)
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node6 := new(ListNode)
	node7 := new(ListNode)
	node8 := new(ListNode)
	node9 := new(ListNode)
	node10 := new(ListNode)
	node11 := new(ListNode)
	node12 := new(ListNode)
	node0.Val = 3
	node1.Val = 0
	node2.Val = 2
	node3.Val = 6
	node4.Val = 8
	node5.Val = 1
	node6.Val = 7
	node7.Val = 9
	node8.Val = 4
	node9.Val = 2
	node10.Val = 5
	node11.Val = 5
	node12.Val = 0
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9
	node9.Next = node10
	node10.Next = node11
	node11.Next = node12
	spiralMatrix(m, n, node0)
}

type pair struct {
	x int
	y int
}

var direction = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = -1
		}
	}
	for x, y, dir := 0, 0, 0; head != nil; head = head.Next {
		result[x][y] = head.Val
		d := direction[dir&3]
		xx, yy := x+d.x, y+d.y
		if xx < 0 || xx >= m || yy < 0 || yy >= n || result[xx][yy] != -1 {
			dir++
			d = direction[dir&3]
		}
		x += d.x
		y += d.y
	}
	return result
}
