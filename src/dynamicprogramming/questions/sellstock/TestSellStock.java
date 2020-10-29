package dynamicprogramming.questions.sellstock;

public class TestSellStock {

    public static void main(String[] args) {
        case1();
        case2();
        case3();
    }

    private static void case1() {
        int[] case1 = new int[]{7, 1, 5, 3, 6, 4};
        int result = (new SellStock()).maxProfit(case1);
        System.out.println(result);
    }

    private static void case2() {
        int[] case1 = new int[]{7,6,4,3,1};
        int result = (new SellStock()).maxProfit(case1);
        System.out.println(result);
    }

    private static void case3() {
        int[] case1 = new int[]{1,2};
        int result = (new SellStock()).maxProfit(case1);
        System.out.println(result);
    }
}
