package array.questions.stocksell;

public class StockSellTest {

    public static void main(String[] args) {
        int[] price1 = {7, 1, 5, 3, 6, 4};
        int[] price2 = {1, 2, 3, 4, 5};
        int[] price3 = {7, 6, 4, 3, 1};

        test(price1, 7);
        test(price2, 4);
        test(price3, 0);

        // See! Only Java can pass null to the function, while Kotlin/Go we can't do that...
        test(null, 0);

        testKt(price1, 7);
        testKt(price2, 4);
        testKt(price3, 0);
    }

    private static void test(int[] prices, int expect) {
        StockSell ss = new StockSell();
        int result = ss.maxProfit(prices);

        System.out.println("test: expected - " + expect + ", result - " + result);
    }

    private static void testKt(int[] prices, int expect) {
        StockSellKt ssKt = new StockSellKt();
        int result = ssKt.maxProfit(prices);

        System.out.println("testkt: expected - " + expect + ", result - " + result);
    }
}
