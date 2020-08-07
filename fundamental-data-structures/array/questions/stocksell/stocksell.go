package questions

// StockSell return max profit
// - if using []int as parameter, then we can't pass array to this function
// - if using [5]int as parameter, then we have limitation here (type-strong)
func StockSell(price []int) int {
	result := 0
	len := len(price)
	if len > 0 {
		for i := 0; i < len-1; i++ {
			gap := price[i+1] - price[i]
			if gap > 0 {
				result += gap
			}
		}

		return result
	}
	return 0
}
