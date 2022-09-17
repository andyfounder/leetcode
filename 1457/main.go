package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	node := new(TreeNode)
	node10 := new(TreeNode)
	node11 := new(TreeNode)
	node20 := new(TreeNode)
	node21 := new(TreeNode)
	node23 := new(TreeNode)
	node.Val = 2
	node10.Val = 3
	node11.Val = 1
	node20.Val = 3
	node21.Val = 1
	node23.Val = 1
	node.Left = node10
	node.Right = node11
	node10.Left = node20
	node10.Right = node21
	node11.Right = node23
	return preorder(root, 0)
}

func preorder(root *TreeNode, cnt int) int {
	if root == nil {
		return 0
	}
	cnt ^= 1 << root.Val
	if root.Left == nil && root.Right == nil && cnt&(cnt-1) == 0 {
		return 1
	}
	return preorder(root.Left, cnt) + preorder(root.Right, cnt)
}
