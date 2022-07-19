package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var node0 TreeNode
	var node10 TreeNode
	var node20 TreeNode
	var node21 TreeNode
	node0.Val = 0
	node0.Left = &node10
	node0.Right = nil
	node10.Val = 0
	node10.Left = &node20
	node10.Right = &node21
	node20.Val = 0
	node20.Left = nil
	node20.Right = nil
	node21.Val = 0
	node21.Left = nil
	node21.Right = nil
	fmt.Println(minCameraCover(&node0))
}

func minCameraCover(root *TreeNode) int {
	return 0
}
