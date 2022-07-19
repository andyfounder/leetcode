package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var listNode1 ListNode
	var listNode2 ListNode
	var listNode3 ListNode
	listNode1.Val = 1
	listNode1.Next = &listNode2
	listNode2.Val = 1
	listNode2.Next = &listNode3
	listNode3.Val = 2
	listNode3.Next = nil
	deleteDuplicates(&listNode1)
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	value := head.Val
	current := head
	temp := head
	for current.Next != nil {
		if current.Next.Val == value {
			current = current.Next
		} else {
			current = current.Next
			value = current.Val
			temp.Next = current
			temp = temp.Next
		}
	}
	temp.Next = nil
	return head
}
