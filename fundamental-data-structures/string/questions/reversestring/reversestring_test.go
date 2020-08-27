package reversestring

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	case1 := []byte("hello")
	fmt.Println(case1)
	ReverseString(case1)

	case2 := []byte("Hannah")
	fmt.Println(case2)
	ReverseString(case2)

	fmt.Println(case1)
	ReverseString1(case1)

	fmt.Println(case2)
	ReverseString1(case2)

	fmt.Println()
	fmt.Println(case1)
	ReverseString2(case1)

	fmt.Println(case2)
	ReverseString2(case2)
}
