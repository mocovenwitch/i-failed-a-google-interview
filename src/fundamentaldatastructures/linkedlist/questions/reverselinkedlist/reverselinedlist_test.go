package reverselinkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	case15 := ListNode{5, nil}
	case14 := ListNode{4, &case15}
	case13 := ListNode{3, &case14}
	case12 := ListNode{2, &case13}
	case11 := ListNode{1, &case12}

	result1 := reverseList(&case11)
	fmt.Println(result1)

	var case2 *ListNode = nil
	result2 := reverseList(case2)
	fmt.Println(result2)

}
