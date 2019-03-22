package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	i1 := listNodeToArr(*l1)
	i2 := listNodeToArr(*l2)
	arr := addTwoArr(i1, i2)
	tmp := arrToListNode(arr)
	return &tmp
}

func listNodeToArr(l ListNode) []int {
	tmp := []int{}
	a := 1
	for {
		tmp = append(tmp, l.Val)
		if l.Next == nil {
			break
		} else {
			l = *l.Next
			a = a * 10
		}
	}
	return tmp

}

func arrToListNode(i []int) ListNode {

	l := ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := &l
	for in, a := range i {
		(*tmp).Val = a
		tmp1 := ListNode{
			Val:  0,
			Next: nil,
		}
		if in == len(i)-1 {
			break
		}
		(*tmp).Next = &tmp1
		tmp = &tmp1
	}
	return l
}

func addTwoArr(a, b []int) []int {

	c := []int{}
	if len(a) > len(b) {
		c = b
		b = a
		a = c
	}
	b = append(b, 0)
	for i := 0; i < len(a); i++ {
		b[i] = (b[i] + a[i])
	}

	for i := 0; i < len(b)-1; i++ {
		b[i+1] = b[i+1] + (b[i] / 10)
		b[i] = b[i] % 10
	}

	if b[len(b)-1] == 0 {
		return b[:len(b)-1]
	} else {
		return b
	}
}

func main() {
	//	a := ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 4,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: nil,
	//			},
	//		},
	//	}
	//	b := ListNode{
	//		Val: 5,
	//		Next: &ListNode{
	//			Val: 6,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: nil,
	//			},
	//		},
	//	}

	a := ListNode{
		Val:  1,
		Next: nil,
	}
	b := ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	fmt.Println(addTwoNumbers(&a, &b))
}
