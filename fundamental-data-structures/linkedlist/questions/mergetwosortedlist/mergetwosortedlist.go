package mergetwosortedlist

// worth optimise...
// Runtime: 4 ms
// Memory Usage: 2.8 MB
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

// Runtime: 4 ms
// Memory Usage: 2.6 MB
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	// protect from nil
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var h *ListNode = nil
	var c *ListNode = nil
	for l1 != nil && l2 != nil {
		if c == nil {
			if l1.Val <= l2.Val {
				c = l1
				l1 = l1.Next
			} else {
				c = l2
				l2 = l2.Next
			}
			h = c
		} else {
			if l1.Val <= l2.Val {
				c.Next = l1
				l1 = l1.Next
			} else {
				c.Next = l2
				l2 = l2.Next
			}
			c = c.Next
		}
	}

	if l1 != nil {
		c.Next = l1
	}
	if l2 != nil {
		c.Next = l2
	}

	return h

}

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 4 ms
// Memory Usage: 2.6 MB
// the same time and space complexity
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// protect from nil
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// go with the rest
	var h *ListNode
	var p *ListNode
	var c *ListNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			c = l1
			l1 = l1.Next
		} else {
			c = l2
			l2 = l2.Next
		}

		if h == nil {
			h = c
		}
		if p != nil {
			p.Next = c
		}
		p = c
	}

	// if something left
	if l1 == nil {
		p.Next = l2
	}
	if l2 == nil {
		p.Next = l1
	}
	return h
}
