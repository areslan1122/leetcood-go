package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	headf := head
	head = head.Next
	headf.Next = nil
	n := &ListNode{}
	for {
		n = head.Next
		head.Next = headf
		if n == nil {
			break
		}
		headf = head
		head = n
	}

	return head
}
