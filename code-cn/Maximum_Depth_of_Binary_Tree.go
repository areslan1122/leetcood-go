package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	depth := 1
	maxDepthHelper(root, 1, &depth)

	return depth
}

func maxDepthHelper(root *TreeNode, dep int, depth *int) {

	dep += 1

	if dep > *depth {
		*depth = dep
	}

	if root.Right != nil {
		maxDepthHelper(root.Right, dep, depth)
	}
	if root.Left != nil {
		maxDepthHelper(root.Left, dep, depth)
	}
}
