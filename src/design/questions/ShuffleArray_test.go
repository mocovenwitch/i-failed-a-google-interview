package questions

import (
	"fmt"
	"testing"
)

func TestShuffle(test *testing.T) {
	i := []int{1, 2, 3}
	s := Constructor(i)

	fmt.Println(s.Orignal)
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
}
