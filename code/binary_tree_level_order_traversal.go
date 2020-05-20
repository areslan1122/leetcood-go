package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

	ans := [][]int{}
	levelOrderHelper(&ans, root, 0)

	return ans
}

func levelOrderHelper(arr *[][]int, root *TreeNode, level int) {

	if root != nil {
		if len(*arr) <= level {
			*arr = append(*arr, []int{})
		}
		tmp := *arr
		tmp[level] = append(tmp[level], root.Val)

		if root.Left != nil {
			levelOrderHelper(arr, root.Left, level+1)
		}
		if root.Right != nil {
			levelOrderHelper(arr, root.Right, level+1)
		}
	}

}

func main() {

}
