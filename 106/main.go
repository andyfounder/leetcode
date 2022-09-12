package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	root := new(TreeNode)
	if len(inorder) == 0 {
		return nil
	}
	root.Val = postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	for i, in := range inorder {
		if in == root.Val {
			root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
			root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):])
		}
	}
	return root
}
