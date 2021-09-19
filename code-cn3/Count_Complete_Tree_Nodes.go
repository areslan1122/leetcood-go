package main


func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left, leftD := checkBT(root.Left)
	_, rightD := checkBT(root.Right)

	if left {
		return powerf2(leftD) + countNodes(root.Right)
	} else {
		return powerf2(rightD) + countNodes(root.Left)
	}

}


func powerf2(x int) int {
	if x == 0 {
		return 1
	}
	return 2 * powerf2(x-1)
}

func checkBT(root *TreeNode) (bool, int) {

	leftTree := root
	left, right := 0,0

	for leftTree != nil {
		left++
		leftTree = leftTree.Left
	}

	for root != nil {
		right++
		root = root.Right
	}

	if left == right {
		return true, left
	} else {
		return false, left
	}
}
