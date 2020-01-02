package sort_list

/*
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

Input: 4->2->1->3
Output: 1->2->3->4
Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dumyHead := new(ListNode)
	dumyHead.Next = head
	tmp := dumyHead

	interval := 1
	for {
		head = dumyHead.Next
		tmp = dumyHead
		for head != nil {
			first := head
			head = cut(head, interval)
			if head == nil {
				return dumyHead.Next
			}
			second := head
			head = cut(head, interval)
			tmp.Next = mergeSortedLists(first, second)
			for ; tmp.Next != nil; tmp = tmp.Next {
			}
		}
		interval *= 2
	}
}

func cut(head *ListNode, n int) *ListNode {
	p := head
	for p != nil && n > 0 {
		p = p.Next
		n--
	}
	if p == nil {
		return nil
	}
	r := p.Next
	p.Next = nil
	return r
}

//
func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return divideAndSort(head)
}

func divideAndSort(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	p := head.Next
	q := head
	for p != nil && p.Next != nil {
		p = p.Next.Next
		q = q.Next
	}
	p = q.Next
	q.Next = nil
	head = divideAndSort(head)
	p = divideAndSort(p)
	return mergeSortedLists(head, p)
}

func mergeSortedLists(first, second *ListNode) *ListNode {
	dumyHead := new(ListNode)
	p := dumyHead
	for first != nil && second != nil {
		if first.Val <= second.Val {
			p.Next = first
			first = first.Next
		} else {
			p.Next = second
			second = second.Next
		}
		p = p.Next
	}

	if first != nil {
		p.Next = first
	} else if second != nil {
		p.Next = second
	}
	return dumyHead.Next
}
