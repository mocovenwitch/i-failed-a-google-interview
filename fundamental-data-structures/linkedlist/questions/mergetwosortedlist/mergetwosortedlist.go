package mergetwosortedlist

// worth optimise...
func mergeTwoLists0(l1 *ListNode, l2 *ListNode) *ListNode {

	// protect from nil
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// merge l2 to l1
	c1 := l1
	var p *ListNode = nil
	for l2 != nil {
		if c1 == nil {
			// if finished loop in l1, combine l2 to l1 tail, and return l1
			p.Next = l2
			return l1
		}
		if l2.Val <= c1.Val {
			// c1 stay the same
			// p should move
			temp := l2.Next
			l2.Next = c1
			if p != nil {
				p.Next = l2
				p = l2
			} else {
				p = l2
				l1 = p
			}
			l2 = temp

		} else {
			// assign c1 to p
			// c1 moves to the next one
			p = c1
			c1 = c1.Next
		}
	}

	return l1
}

func mergeTwoLists0(l1 *ListNode, l2 *ListNode) *ListNode {
}

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
