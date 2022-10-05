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
	node20 := new(TreeNode)
	node22 := new(TreeNode)
	node23 := new(TreeNode)
	node30 := new(TreeNode)
	node31 := new(TreeNode)
	node37 := new(TreeNode)
	node.Val = 6
	node10.Val = 4
	node11.Val = 8
	node20.Val = 11
	node22.Val = 13
	node23.Val = 4
	node30.Val = 7
	node31.Val = 2
	node37.Val = 1
	node.Left = node10
	node.Right = node11
	node10.Left = node20
	node11.Left = node22
	node11.Right = node23
	node20.Left = node30
	node20.Right = node31
	node23.Right = node37
	targetSum := 23
	hasPathSum(node, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
