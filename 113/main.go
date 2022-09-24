package main

func main() {
	node := new(TreeNode)
	node10 := new(TreeNode)
	node11 := new(TreeNode)
	node20 := new(TreeNode)
	node22 := new(TreeNode)
	node23 := new(TreeNode)
	node30 := new(TreeNode)
	node31 := new(TreeNode)
	node36 := new(TreeNode)
	node37 := new(TreeNode)
	node.Left = node10
	node.Right = node11
	node.Val = 5
	node10.Left = node20
	node10.Val = 4
	node11.Left = node22
	node11.Right = node23
	node11.Val = 8
	node20.Left = node30
	node20.Right = node31
	node20.Val = 11
	node22.Val = 13
	node23.Left = node36
	node23.Right = node37
	node23.Val = 4
	node30.Val = 7
	node31.Val = 2
	node36.Val = 5
	node37.Val = 1
	targetSum := 23
	pathSum(node, targetSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	stack := []int{}
	res = dfs(root, targetSum, res, stack)
	return res
}

func dfs(root *TreeNode, targetSum int, res [][]int, stack []int) [][]int {
	if root == nil {
		return res
	}
	targetSum -= root.Val
	stack = append(stack, root.Val)
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		res = append(res, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}
	res = dfs(root.Left, targetSum, res, stack)
	res = dfs(root.Right, targetSum, res, stack)
	return res
}
