/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sort_list

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
	p := dumyHead

	interval := 1
	for {
		head = dumyHead.Next
		p = dumyHead
		for head != nil {
			first := head
			head = divide(head, interval)
			if head == nil {
				return dumyHead.Next
			}
			second := head
			head = divide(head, interval)
			p.Next = mergeSortedLists(first, second)
			for ; p.Next != nil; p = p.Next {
			}
		}
		interval *= 2
	}
}

/*
divide list to 2 parts, and returns the head of the second part
if n is too big and list is tool short, returns nil
*/
func divide(head *ListNode, n int) *ListNode {
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
