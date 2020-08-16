package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	rightDeep := getDeep(root.Right, 0)
	leftDeep := getDeep(root.Left, 0)

	if rightDeep-leftDeep > 1 || leftDeep-rightDeep > 1 {
		return false
	}

	return isBalanced(root.Right) && isBalanced(root.Left)
}

func getDeep(root *TreeNode, deep int) int {

	if root == nil {
		return deep
	}

	rightDeep := getDeep(root.Right, deep+1)
	leftDeep := getDeep(root.Left, deep+1)

	if rightDeep > leftDeep {
		return rightDeep
	} else {
		return leftDeep
	}
}
