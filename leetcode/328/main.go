package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{Val: 0, Next: nil}
	odd := oddHead
	evenHead := &ListNode{Val: 0, Next: nil}
	even := evenHead

	count := 1
	for head != nil {
		if count%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		count++
	}
	even.Next = nil
	odd.Next = evenHead.Next
	return oddHead.Next
}

func printLinkList(head *ListNode, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	var head *ListNode = &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: nil,
				Val:  2,
			},
			Val: 3,
		},
		Val: 1,
	}
	res := oddEvenList(head)
	printLinkList(res, 3)

}
