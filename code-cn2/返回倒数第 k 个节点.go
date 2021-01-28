package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {

	arr := []int{}

	for {
		if head == nil {
			break
		}
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr[len(arr)-k+1]
}
