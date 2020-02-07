package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	return twoTreeIsSummetric(root, root)
}

func twoTreeIsSummetric(treeA, treeB *TreeNode) bool {

	if treeA == nil && treeB == nil {
		return true
	}

	if treeA == nil || treeB == nil {
		return false
	}

	return treeA.Val == treeB.Val &&
		twoTreeIsSummetric(treeA.Right, treeB.Left) &&
		twoTreeIsSummetric(treeA.Left, treeB.Right)
}
