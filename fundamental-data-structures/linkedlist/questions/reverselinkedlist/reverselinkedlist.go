package reverselinkedlist

/**
 * Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
 * Memory Usage: 2.6 MB, less than 82.44% of Go online submissions for Reverse Linked List.
 */
func reverseList(head *ListNode) *ListNode {
	// this nil test case caught me at first time submit
	if head == nil {
		return head
	}
	var p *ListNode = nil
	var n *ListNode = nil
	for head.Next != nil {
		n = head.Next
		head.Next = p
		p = head
		head = n
	}
	head.Next = p
	return head
}

// ListNode is
// Definition for singly-linked list.
//  type ListNode struct {
//     Val int
//     Next *ListNode
//  }
type ListNode struct {
	Val  int
	Next *ListNode
}
