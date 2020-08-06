package array.questions;

public class StockSellTest {

    public static void main(String[] args) {
        int[] price1 = {7, 1, 5, 3, 6, 4};
        int[] price2 = {1, 2, 3, 4, 5};
        int[] price3 = {7, 6, 4, 3, 1};

        test(price1, 7);
        test(price2, 4);
        test(price3, 0);
    }

    private static void test(int[] prices, int expect) {
        StockSell ss = new StockSell();
        int result = ss.maxProfit(prices);

        System.out.println("test: expected - " + expect + ", result - " + result);
    }
}
