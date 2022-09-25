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
	targetSum := 22
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

// type pair struct {
// 	node *TreeNode
// 	sum  int
// }

// func pathSum(root *TreeNode, targetSum int) [][]int {
// 	res := [][]int{}
// 	if root == nil {
// 		return res
// 	}
// 	tree := map[*TreeNode]*TreeNode{}
// 	queue := []pair{{root, targetSum}}
// 	getPath := func(node *TreeNode) []int {
// 		path := []int{}
// 		for ; node != nil; node = tree[node] {
// 			path = append(path, node.Val)
// 		}
// 		for i, j := 0, len(path)-1; i < j; i++ {
// 			path[i], path[j] = path[j], path[i]
// 			j--
// 		}
// 		return path
// 	}
// 	for len(queue) > 0 {
// 		q := queue[0]
// 		node := q.node
// 		sum := q.sum - node.Val
// 		queue = queue[1:]
// 		if node.Left == nil && node.Right == nil && sum == 0 {
// 			res = append(res, getPath(node))
// 		} else {
// 			if node.Left != nil {
// 				queue = append(queue, pair{node.Left, sum})
// 				tree[node.Left] = node
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, pair{node.Right, sum})
// 				tree[node.Right] = node
// 			}
// 		}
// 	}
// 	return res
// }
