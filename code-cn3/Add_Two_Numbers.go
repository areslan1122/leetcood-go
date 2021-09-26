package main

type ListNode struct {
    Val int
    Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }

    copy := l1
    res := l1

	for {
        l1.Val += l2.Val
        if l1.Next == nil {
            l1.Next = l2.Next
            break
        } else if l2.Next == nil {
            break
        } else {
            l1 = l1.Next
            l2 = l2.Next
        }
    }

    for copy != nil{

        if copy.Val >= 10 {
            copy.Val %= 10
            if copy.Next == nil {
                copy.Next = &ListNode{
                    Val: 1,
                    Next: nil,
                }
            } else {
                copy.Next.Val += 1
            }
        }

        copy = copy.Next
    }

    return res
}
