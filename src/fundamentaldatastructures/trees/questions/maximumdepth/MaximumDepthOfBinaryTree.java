package trees.questions.maximumdepth;


/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
class MaximumDepthOfBinaryTree {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    /**
     * time: beat 100%
     * memory: beat 73.58%
     */
    public int maxDepth0(TreeNode root) {
        int dep = 0;

        if (root != null) {
            dep++;
            int leftDep = 0, rightDep = 0;
            if (root.left != null) {
                leftDep = maxDepth(root.left);
            }
            if (root.right != null) {
                rightDep = maxDepth(root.right);
            }
            dep += Math.max(leftDep, rightDep);
        }
        return dep;
    }

    /**
     * time: beat 100%
     * memory: beat 51.55% ???
     */
    public int maxDepth(TreeNode root) {
        if (root == null) return 0;
        return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
    }
}