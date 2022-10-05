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
	node23 := new(TreeNode)
	node32 := new(TreeNode)
	node33 := new(TreeNode)
	node.Val = 3
	node10.Val = 5
	node11.Val = 1
	node20.Val = 6
	node21.Val = 2
	node22.Val = 0
	node23.Val = 8
	node32.Val = 7
	node33.Val = 4
	node.Left = node10
	node.Right = node11
	node10.Left = node20
	node10.Right = node21
	node11.Left = node22
	node11.Right = node23
	node21.Right = node32
	node21.Right = node33
	subtreeWithAllDeepest(node)
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parents := make(map[*TreeNode]*TreeNode, 0)
	queue := []*TreeNode{root}
	temp := queue
	for len(queue) > 0 {
		length := len(queue)
		temp = queue
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				parents[queue[i].Left] = queue[i]
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				parents[queue[i].Right] = queue[i]
			}
		}
		queue = queue[length:]
	}
	find := func() bool {
		for i := 0; i < len(temp); i++ {
			if temp[i] != temp[0] {
				return false
			}
		}
		return true
	}
	for find() != true {
		for i := 0; i < len(temp); i++ {
			temp[i] = parents[temp[i]]
		}
	}
	return temp[0]
}
