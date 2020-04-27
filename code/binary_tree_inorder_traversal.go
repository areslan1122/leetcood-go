package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {

	if root != nil {
		return []int{}
	}

	l := inorderTraversal(root.Left)
	r := inorderTraversal(root.Right)
	l = append(l, root.Val)
	l = append(l, r...)

	return l
}
