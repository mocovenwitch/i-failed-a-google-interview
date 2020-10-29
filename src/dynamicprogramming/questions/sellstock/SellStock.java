package dynamicprogramming.questions.sellstock;

public class SellStock {

    /**
     * O(n^2), 5.04% time, 8.00% memory
     */
    public int maxProfit_1(int[] prices) {
        // protect from 1 and less stock
        if (prices.length <= 1) {
            return 0;
        }

        int max = 0;
        for(int i = 0; i < prices.length-1; i++) {
            for (int j = i + 1; j < prices.length; j++) {
                if ((prices[i] < prices[j]) && (prices[j] - prices[i] > max)) {
                    max = prices[j] - prices[i];
                }
            }
        }

        return max;
    }

    /**
     * O(n), beat 98.38% time, less than 8.00% memory
     */
    public int maxProfit(int[] prices) {
        // protect from 1 and less stock
        if (prices.length <= 1) {
            return 0;
        }

        int min = prices[0];
        int result = 0;
        for(int i = 0; i < prices.length; i++) {
            if (prices[i] < min) {
                min = prices[i];
            }
            if (prices[i] - min > result) {
                result = prices[i] - min;
            }
        }

        return result;
    }
}
