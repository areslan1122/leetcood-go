package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {

	list := getAllNode(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}

}

func getAllNode(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}

	if root != nil {

		list = append(list, root)
		list = append(list, getAllNode(root.Left)...)
		list = append(list, getAllNode(root.Right)...)
	}

	return list
}
