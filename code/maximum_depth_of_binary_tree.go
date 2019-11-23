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

	if root == nil {
		return 0
	}

	root.Val = 1
	updateDepth(root, root)

	return root.Val

}

func updateDepth(root *TreeNode, subRoot *TreeNode) {

	if subRoot.Val > root.Val {
		root.Val = subRoot.Val
	}
	if subRoot.Left != nil {
		subRoot.Left.Val = subRoot.Val + 1
		updateDepth(root, subRoot.Left)
	}
	if subRoot.Right != nil {
		subRoot.Right.Val = subRoot.Val + 1
		updateDepth(root, subRoot.Right)
	}
}
