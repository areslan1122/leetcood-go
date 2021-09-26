package main

func minDepth(root *TreeNode) int {

}

func minDepthHelper(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := minDepthHelper(root.Left)
	right := minDepthHelper(root.Right)

	if left < right {
		return left + 1
	} else {
		return right + 1
	}

}



