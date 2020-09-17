package trees.questions.maximumdepth;

import trees.questions.maximumdepth.MaximumDepthOfBinaryTree.TreeNode;

public class TestMaximumDepth {
    public static void main(String[] ars) {
        //[3,9,20,null,null,15,7]

        TreeNode sevenNode = new TreeNode(7, null, null);
        TreeNode fifteenNode = new TreeNode(15, null, null);
        TreeNode twentyNode = new TreeNode(20, fifteenNode, sevenNode);
        TreeNode nineNode = new TreeNode(9, null, null);
        TreeNode threeNode = new TreeNode(3, nineNode, twentyNode);

        MaximumDepthOfBinaryTree trees = new MaximumDepthOfBinaryTree();
        int result = trees.maxDepth(threeNode);
        System.out.println(result);
    }
}
