package main

func main() {
	node := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	n := 2
	removeNthFromEnd(node, n)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp, length := head, 0
	for temp != nil {
		length++
		temp = temp.Next
	}
	if length-n == 0 {
		if head.Next != nil {
			head = head.Next
		} else {
			head = nil
		}
		return head
	}
	temp = head
	for i := 1; i < length-n; i++ {
		temp = temp.Next
	}
	if temp.Next.Next != nil {
		temp.Next = temp.Next.Next
	} else {
		temp.Next = nil
	}
	return head
}
