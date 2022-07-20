package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := new(TreeNode)
	node11 := new(TreeNode)
	node12 := new(TreeNode)
	node23 := new(TreeNode)
	node24 := new(TreeNode)
	node.Val = 3
	node.Left = node11
	node.Right = node12
	node11.Val = 9
	node12.Val = 20
	node12.Left = node23
	node12.Right = node24
	node23.Val = 15
	node24.Val = 7
	levelOrder(node)
}

//BFS
// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return [][]int{}
// 	}
// 	res := make([][]int, 0)
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0 {
// 		l := len(queue)
// 		tmp := make([]int, 0, l)
// 		for i := 0; i < l; i++ {
// 			if queue[i].Left != nil {
// 				queue = append(queue, queue[i].Left)
// 			}
// 			if queue[i].Right != nil {
// 				queue = append(queue, queue[i].Right)
// 			}
// 			tmp = append(tmp, queue[i].Val)
// 		}
// 		queue = queue[l:]
// 		res = append(res, tmp)
// 	}
// 	return res
// }

//DFS
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}
