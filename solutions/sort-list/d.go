/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sort_list

import "sort"

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

/*
 归并排序：自顶向下，使用递归
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	right := divide(head)
	return merge(sortList(head), sortList(right))
}
func divide(head *ListNode) *ListNode {
	fast := head.Next // 最终返回的是中间节点的下一个节点；写成 fast := head.Next.Next也行，最终返回的是中间节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next
	}
	fast = head.Next
	head.Next = nil
	return fast
}
func merge(first, second *ListNode) *ListNode {
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
	p = dumyHead.Next
	dumyHead.Next, dumyHead = nil, nil
	return p
}

/*
快排，自顶向下，使用递归
选择一个标准值，将比它大的放在一个链表中，比它小的放在一个链表中，和它一样大的，放在另一个链表中。
然后针对小的和大的链表，继续排序。最终将三个链表按照小、相等、大进行连接。
*/
func sortList2(head *ListNode) *ListNode {
	return quickSort(head)
}
func quickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lowDumy, midDumy, highDumy := new(ListNode), new(ListNode), new(ListNode)
	low, mid, high := lowDumy, midDumy, highDumy
	val := head.Val
	for p := head; p != nil; p = p.Next {
		if p.Val < val {
			low.Next = p
			low = low.Next
		} else if p.Val > val {
			high.Next = p
			high = high.Next
		} else {
			mid.Next = p
			mid = mid.Next
		}
	}
	low.Next, mid.Next, high.Next = nil, nil, nil
	lowDumy.Next = quickSort(lowDumy.Next)
	highDumy.Next = quickSort(highDumy.Next)
	low = lowDumy
	for low.Next != nil {
		low = low.Next
	}
	low.Next = midDumy.Next
	mid.Next = highDumy.Next
	low = lowDumy.Next
	lowDumy, midDumy, highDumy = nil, nil, nil
	return low
}

/*
可以用一个数组，装入链表所有节点，然后用标准库对数组排序即可
时间复杂度是O(nlogn), 空间复杂度O(n)
*/
func sortList9(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var arr []*ListNode
	for ; head != nil; head = head.Next {
		arr = append(arr, head)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	for i, v := range arr {
		if i == len(arr)-1 {
			v.Next = nil
		} else {
			v.Next = arr[i+1]
		}
	}
	return arr[0]
}
