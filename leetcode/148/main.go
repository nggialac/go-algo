package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	min := head.Val
	var tmp *ListNode = head.Next
	for head != nil {
		for tmp != nil {
			if tmp.Val < min {
				min = head.Val
			}
			tmp = tmp.Next
		}
		head.Val = min
		head = head.Next
		tmp = head
		min = head.Val
	}
	return head
}

func display(head *ListNode, n int) {
	// for head != nil {
	// 	fmt.Println(head.Val)
	// 	head = head.Next
	// }
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
	sortList(head)
	// display(res, 2)
}
