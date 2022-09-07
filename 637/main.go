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
	node22 := new(TreeNode)
	node23 := new(TreeNode)
	node.Val = 3
	node.Left = node10
	node.Right = node11
	node10.Val = 9
	node11.Val = 20
	node11.Left = node22
	node11.Right = node23
	node22.Val = 15
	node23.Val = 7
	averageOfLevels(node)
}
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}
	total := 0.0
	num := 0
	queue := []*TreeNode{root, nil}
	for len(queue) > 1 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			queue = append(queue, nil)
			res = append(res, total/float64(num))
			total = 0
			num = 0
		} else {
			total += float64(node.Val)
			num++
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	res = append(res, total/float64(num))
	return res
}
