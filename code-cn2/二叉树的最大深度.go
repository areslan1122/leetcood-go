package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	ans := 0
	ans = maxDepthHelper(root, ans)
	return ans
}

func maxDepthHelper(root *TreeNode, depth int) int {

	if root == nil {
		return depth
	}
	left := maxDepthHelper(root.Left, depth+1)
	right := maxDepthHelper(root.Right, depth+1)

	if left > right {
		depth = left
	} else {
		depth = right
	}
	return depth
}
