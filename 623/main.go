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
	node.Val = 4
	node10.Val = 2
	node11.Val = 6
	node20.Val = 3
	node21.Val = 1
	node22.Val = 5
	node.Left = node10
	node.Right = node11
	node10.Left = node20
	node10.Right = node21
	node11.Left = node22
	val := 1
	depth := 2
	addOneRow(node, val, depth)
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	if depth == 2 {
		root.Left = &TreeNode{val, root.Left, nil}
		root.Right = &TreeNode{val, nil, root.Right}
	} else {
		if root.Left != nil {
			addOneRow(root.Left, val, depth-1)
		}
		if root.Right != nil {
			addOneRow(root.Right, val, depth-1)
		}
	}
	return root
}

// func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
// 	if depth == 1 {
// 		return &TreeNode{val, root, nil}
// 	}
// 	queue := []*TreeNode{root}
// 	for i := 1; i < depth-1; i++ {
// 		length := len(queue)
// 		for j := 0; j < length; j++ {
// 			node := queue[j]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		queue = queue[length:]
// 	}
// 	for _, node := range queue {
// 		node.Left = &TreeNode{val, node.Left, nil}
// 		node.Right = &TreeNode{val, nil, node.Right}
// 	}
// 	return root
// }
