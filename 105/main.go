package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	array1 := []int{3, 9, 20, 15, 7}
	array2 := []int{9, 3, 15, 20, 7}
	buildTree(array1, array2)
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[0]
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}
