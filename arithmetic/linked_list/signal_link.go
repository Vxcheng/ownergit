package main

// 1. 反转位置m-n的单链表，请使用一趟扫描完成反转
// 例： 1->2->3->4->5  reverse: 1->4->3->2->5

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseElement(head *ListNode, m, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	he := &ListNode{}
	he.Next = head
	cur := he

	for i := 0; i < m-1; i++ {
		cur = cur.Next
	}

	h := cur.Next
	t := h.Next

	var tail *ListNode
	for i := 0; i < n-m; i++ {
		temp := t.Next
		t.Next = h
		if tail == nil {
			tail = h
		}

		h = t
		t = temp
	}

	cur.Next = h
	if tail != nil {
		tail.Next = t
	}
	return he.Next
}
