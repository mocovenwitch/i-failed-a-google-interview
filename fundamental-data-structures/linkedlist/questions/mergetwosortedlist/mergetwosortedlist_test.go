package mergetwosortedlist

import (
	"fmt"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	case0()
	case1()
}

func case0() {
	case14 := ListNode{4, nil}
	case12 := ListNode{2, &case14}
	case11 := ListNode{1, &case12}

	case24 := ListNode{4, nil}
	case23 := ListNode{3, &case24}
	case21 := ListNode{1, &case23}

	result := mergeTwoLists(&case11, &case21)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next

	}
}

func case1() {
	case14 := ListNode{2, nil}
	case12 := ListNode{2, &case14}
	case11 := ListNode{2, &case12}

	case24 := ListNode{1, nil}
	case23 := ListNode{1, &case24}
	case21 := ListNode{1, &case23}

	result := mergeTwoLists(&case11, &case21)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next

	}
}
