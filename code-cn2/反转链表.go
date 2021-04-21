package main

//type ListNode struct {
//    Val int
//    Next *ListNode
//}
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	} else if head == nil || head.Next == nil{
		return head
	}

	one := head
	one = nil
	two := head

	for {
		tmp := two.Next
		two.Next = one
		if tmp == nil {
			break
		}

		one = two
		two = tmp
	}

	return two

}
