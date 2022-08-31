package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := new(ListNode)
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node.Next = node1
	node.Val = 1
	node1.Next = node2
	node1.Val = 2
	node2.Next = node3
	node2.Val = 2
	node3.Val = 1
	fmt.Println(isPalindrome(node))
}
func isPalindrome(head *ListNode) bool {
	slice := make([]int, 0)
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	i, j := 0, len(slice)-1
	for i < j {
		if slice[i] != slice[j] {
			return false
		}
		i++
		j--
	}
	return true
}
