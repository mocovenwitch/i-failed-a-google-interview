package array.questions.stocksell

class StockSellKt {

    fun maxProfit(prices: IntArray): Int {

        // protect from empty
        val len = prices.size
        if (len <= 0) return 0

        // start
        var result = 0
        for (i in 0 until len - 1) {
            val gap = prices[i + 1] - prices[i]
            if (gap > 0) {
                result += gap
            }
        }

        return result
    }
}