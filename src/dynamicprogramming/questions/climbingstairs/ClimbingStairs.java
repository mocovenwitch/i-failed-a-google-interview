package dynamicprogramming.questions.climbingstairs;

public class ClimbingStairs {

    // Time Limit Exceeded
    public int climbStairs0(int n) {
        if (n <= 3) {
            return n;
        } else {
            return climbStairs0(n - 1) + climbStairs0(n - 2);
        }
    }

    // Memorize strategy.
    // memory beat 70%
    private final int[] memo = new int[45];
    public int climbStairs(int n) {
        if (n <= 3) {
            return n;
        } else {
            int first = memo[n - 1] != 0 ? memo[n - 1] : climbStairs(n - 1);
            int second = memo[n - 2] != 0 ? memo[n - 2] : climbStairs(n - 2);
            memo[n-1] = first;
            memo[n-2] = second;
            return first + second;
        }
    }

    // time: 0ms
    // memory: 90.49%
    public int climbStairs2(int n) {
        if (n <= 3) return n;

        int first = 2;
        int second = 3;
        for (int i = 4; i <= n; i++) {
            int temp = first + second;
            first = second;
            second = temp;
        }
        return second;
    }
}
