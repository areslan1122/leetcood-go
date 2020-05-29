package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	ans := &ListNode{}
	tmp := ans

	for {
		if l1 == nil {
			tmp.Next = l2
			break
		} else if l2 == nil {
			tmp.Next = l1
			break
		}

		if l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}

	return ans.Next
}
