package main


func preorderTraversal(root *TreeNode) []int {

	res := []int{}

	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}