package main

func isSymmetric(root *TreeNode) bool {

	if root == nil {return true}

	return isSameTreeHelper(root.Left, root.Right)
}


func isSameTreeHelper(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	return (p.Val == q.Val) && isSameTreeHelper(p.Right, q.Left) && isSameTreeHelper(p.Left, q.Right)
}