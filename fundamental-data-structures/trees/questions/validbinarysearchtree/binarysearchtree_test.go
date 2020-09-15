package validbinarysearchtree

import "testing"

func TestBinaryTree(t *testing.T) {
	test1() // false
	test2() // true
	test3() // true
}

func test1() {
	firstLeft := TreeNode{1, nil, nil}
	first := TreeNode{1, &firstLeft, nil}

	result := isValidBST(&first)
	print(result)
}

func test2() {
	firstLeft := TreeNode{-1, nil, nil}
	first := TreeNode{0, &firstLeft, nil}

	result := isValidBST(&first)
	print(result)
}

func test3() {
	first := TreeNode{2147483647, nil, nil}

	result := isValidBST(&first)
	print(result)
}
