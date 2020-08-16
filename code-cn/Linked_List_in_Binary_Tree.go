package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {

	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	return isSub(head, root) || isSubPath(head, root.Right) || isSubPath(head, root.Left)
}

func isSub(head *ListNode, root *TreeNode) bool {

	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	if head.Val != root.Val {
		return false
	}
	return isSub(head.Next, root.Right) || isSub(head.Next, root.Left)
}
