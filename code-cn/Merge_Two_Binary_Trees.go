package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return t1
	} else if t1 == nil {
		t1 = t2
	} else if t2 == nil {
		return t1
	} else {
		t1.Val += t2.Val
	}
	_ = mergeTrees(t1.Left, t2.Left)
	_ = mergeTrees(t1.Right, t2.Right)

	return t1
}
