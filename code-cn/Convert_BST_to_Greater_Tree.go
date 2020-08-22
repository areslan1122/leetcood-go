package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func convertBST(root *TreeNode) *TreeNode {

	var add int
	var helper func(root *TreeNode)

	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Right)
		root.Val += add
		add = root.Val
		helper(root.Left)
	}

	helper(root)
	return root
}
