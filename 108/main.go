package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arrays := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(arrays)
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := new(TreeNode)
	root.Val = nums[len(nums)/2]
	if len(nums)/2 == 0 {
		root.Left = nil
	} else {
		root.Left = sortedArrayToBST(nums[0 : len(nums)/2])
	}
	if len(nums)/2+1 >= len(nums) {
		root.Right = nil
	} else {
		root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	}
	return root
}
