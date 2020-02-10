package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	myMap := make(map[*ListNode]bool)

	for head != nil {
		if myMap[head] {
			return true
		}
		mymap[head] = true
		head = head.Next
	}

	return false
}
