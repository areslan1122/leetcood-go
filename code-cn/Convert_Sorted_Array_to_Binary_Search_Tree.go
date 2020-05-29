package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{}
	buildTree(root, nums)

	return root

}

func buildTree(root *TreeNode, nums []int) {

	mi := (len(nums) - 1) / 2
	fmt.Println("mi: ", mi)
	root.Val = nums[mi]
	fmt.Println("nums: ", nums, "root: ", *root)

	if mi != 0 {
		root.Left = &TreeNode{}
		root.Right = &TreeNode{}
		buildTree(root.Left, nums[:mi])
		buildTree(root.Right, nums[(mi+1):])
	} else if len(nums[(mi+1):]) > 0 {
		root.Right = &TreeNode{}
		buildTree(root.Right, nums[(mi+1):])
	}
}

func main() {

	a := []int{-10, -3, 0, 5, 9}
	fmt.Println(*sortedArrayToBST(a))
}
