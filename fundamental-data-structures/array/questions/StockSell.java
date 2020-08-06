package array.questions;

public class StockSell {

    /**
     * O(n) solution
     */
    public int maxProfit(int[] prices) {
        // protect from null
        if (prices == null) return 0;

        // protect from empty
        int len = prices.length;
        if (len <= 0) return 0;

        // start
        int result = 0;
        for (int i = 0; i < len - 1; i++) {
            if (prices[i] < prices[i + 1]) {
                result += prices[i + 1] - prices[i];
            }
        }

        return result;
    }
}
