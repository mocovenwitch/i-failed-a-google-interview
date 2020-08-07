package questions

import (
	"fmt"
	"testing"
)

func TestStockSell(t *testing.T) {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{7, 6, 4, 3, 1}

	result1 := StockSell(prices1)
	result2 := StockSell(prices2)
	result3 := StockSell(prices3)

	fmt.Println("result ", result1, "expected 7")
	fmt.Println("result ", result2, "expected 4")
	fmt.Println("result ", result3, "expected 0")
}
