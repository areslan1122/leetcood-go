package main

func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return sumOFLeftLeavesHelper(root, false)
}

func sumOFLeftLeavesHelper(root *TreeNode, site bool) int {

	tmp := false

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		tmp = true
	}

	if site && tmp {
		return root.Val + sumOFLeftLeavesHelper(root.Left, true) + sumOFLeftLeavesHelper(root.Right, false)
	} else {
		return sumOFLeftLeavesHelper(root.Left, true) + sumOFLeftLeavesHelper(root.Right, false)
	}
}
