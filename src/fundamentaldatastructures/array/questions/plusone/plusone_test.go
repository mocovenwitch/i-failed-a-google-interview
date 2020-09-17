package plusone

import (
	"fmt"
	"testing"
)

func TestStockSell(t *testing.T) {
	test1 := []int{1, 2, 9}
	fmt.Println(PlusOne(test1))

	test2 := []int{9, 9, 9}
	fmt.Println(PlusOne(test2))

	test3 := []int{9, 0, 9}
	fmt.Println(PlusOne(test3))
}
