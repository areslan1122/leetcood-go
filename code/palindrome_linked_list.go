package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	tmp := []int{}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i] != tmp[len(tmp)-1-i] {
			return false
		}
	}
	return true
}
