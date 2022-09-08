package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := new(TreeNode)
	node11 := new(TreeNode)
	node22 := new(TreeNode)
	node.Val = 1
	node11.Val = 2
	node22.Val = 3
	node.Right = node11
	node11.Left = node22
	inorderTraversal(node)
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	find(root, &res)
	return res
}
func find(root *TreeNode, res *[]int) {
	if root != nil {
		find(root.Left, res)
		*res = append(*res, root.Val)
		find(root.Right, res)
	}
}
