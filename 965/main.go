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
	node21 := new(TreeNode)
	node22 := new(TreeNode)
	node.Val = 1
	node10.Val = 1
	node11.Val = 1
	node20.Val = 1
	node21.Val = 1
	node22.Val = 1
	node.Left = node10
	node.Right = node11
	node10.Left = node20
	node10.Right = node21
	node11.Left = node22
	isUnivalTree(node)
}

func isUnivalTree(root *TreeNode) bool {
	res := true
	if root != nil {
		value := root.Val
		if root.Left != nil {
			if value != root.Left.Val {
				res = false
				return res
			}
		}
		res1 := isUnivalTree(root.Left)
		res = res && res1
		if root.Right != nil {
			if value != root.Right.Val {
				res = false
				return res
			}
		}
		res2 := isUnivalTree(root.Right)
		res = res && res2
	}
	return res
}
