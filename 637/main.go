package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := new(TreeNode)
	node10 := new(TreeNode)
	node11 := new(TreeNode)
	node22 := new(TreeNode)
	node23 := new(TreeNode)
	node.Val = 3
	node.Left = node10
	node.Right = node11
	node10.Val = 9
	node11.Val = 20
	node11.Left = node22
	node11.Right = node23
	node22.Val = 15
	node23.Val = 7
	averageOfLevels(node)
}
func averageOfLevels(root *TreeNode) []float64 {
	array := make([]int, 0)
	for root != nil {
		
	}
}
