package removeNthFromEnd

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head
	p := preHead
	q := preHead
	for i := 0; i <= n; i ++ {
		q = q.Next
	}

	if q == nil {
		return head.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}

	// delete node p.Next
	p.Next = p.Next.Next

	return head
}
