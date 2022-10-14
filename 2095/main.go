/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	node1, node2 := head, head
	var node3 *ListNode
	for node1 != nil && node1.Next != nil {
		node1 = node1.Next.Next
		node3 = node2
		node2 = node2.Next
	}
	node3.Next = node2.Next
	return head
}