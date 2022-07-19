package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNodeA0 := new(ListNode)
	listNodeA1 := new(ListNode)
	listNodeA2 := new(ListNode)
	listNodeA0.Val = 1
	listNodeA0.Next = listNodeA1
	listNodeA1.Val = 2
	listNodeA1.Next = listNodeA2
	listNodeA2.Val = 3

	var listNodeB0 *ListNode
	fmt.Println(getIntersectionNode(listNodeA0, listNodeB0))
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
			fmt.Println("111")
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
