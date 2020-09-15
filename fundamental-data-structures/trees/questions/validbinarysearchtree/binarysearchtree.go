package validbinarysearchtree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 * RESULT: 97.12% time; 96.58% memory
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTNode(root, -math.MaxInt64, math.MaxInt64)
}

func isValidBSTNode(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBSTNode(root.Left, min, root.Val) && isValidBSTNode(root.Right, root.Val, max)
}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
