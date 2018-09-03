package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeNode(tree1 *TreeNode, tree2 *TreeNode) {

	tree1.Val += tree2.Val

	if tree1.Left == nil && tree2.Left != nil {
		tree1.Left = tree2.Left
	}
	if tree1.Right == nil && tree2.Right != nil {
		tree1.Right = tree2.Right
	}
	if tree1.Left != nil && tree2.Left != nil {
		mergeNode(tree1.Left, tree2.Left)
	}
	if tree1.Right != nil && tree2.Right != nil {
		mergeNode(tree1.Right, tree2.Right)
	}

}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil && t2 != nil {
		return t2
	}

	if t1 != nil && t2 == nil {
		return t1
	}

	mergeNode(t1, t2)
	return t1
}

func main() {

	a := TreeNode{Val: 3}
	fmt.Println(a)

	b := TreeNode{Val: 1}
	c := TreeNode{Val: 2}
	a = TreeNode{Val: 3, Left: &b, Right: &c}
	d := a

	fmt.Println(a)
	A := mergeTrees(&a, &d)
	fmt.Println(A)

	fmt.Println(a.Right)
	fmt.Println(a.Left)
}
