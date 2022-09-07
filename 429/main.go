package main

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	node := new(Node)
	node10 := new(Node)
	node11 := new(Node)
	node12 := new(Node)
	node20 := new(Node)
	node21 := new(Node)
	node.Val = 1
	node10.Val = 3
	node11.Val = 2
	node12.Val = 4
	node20.Val = 5
	node21.Val = 6
	node.Children = []*Node{node10, node11, node12}
	node10.Children = []*Node{node20, node21}
	levelOrder(node)
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	var temp []int
	if root == nil {
		return res
	}
	queue := []*Node{root, nil}
	for len(queue) > 1 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			queue = append(queue, nil)
			res = append(res, temp)
			temp = []int{}
		} else {
			temp = append(temp, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
	}
	res = append(res, temp)
	return res
}
