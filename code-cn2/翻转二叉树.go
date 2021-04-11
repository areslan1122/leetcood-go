package main

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	_ = invertTree(root.Left)
	_ = invertTree(root.Right)
	tmp := root.Left
	root.Left = root.Right
	root.Right =tmp

	return root
}