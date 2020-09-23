package dynamicprogramming.questions.climbingstairs;

public class TestClimbingStairs {
    public static void main(String[] args) {
        ClimbingStairs climbingStairs = new ClimbingStairs();

        int n3 = climbingStairs.climbStairs(6);
        System.out.println("result:" + n3);

        int n = climbingStairs.climbStairs(7);
        System.out.println("result:" + n);

        int n1 = climbingStairs.climbStairs(8);
        System.out.println("result:" + n1);
    }
}
