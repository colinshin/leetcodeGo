package alindrome_linked_list

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*如果是数组，则很容易判断是否回文，从两端逐渐往中间遍历，每次比较两端是否相等即可，不再赘述

用 O(n) 时间复杂度和 O(1) 空间复杂度解决
思路是将后半段反转后和前半段对比
*/
func isPalindrome1(head *ListNode) bool {
	n := 0 //计算链表长度
	for p := head; p != nil; p = p.Next {
		n++
	}
	// 找到链表中点的下一个位置
	p := head
	for i := 0; i < n/2; i++ {
		p = p.Next
	}
	if n%2 == 1 {
		p = p.Next
	}
	// 翻转后半段链表
	p = reverse(p)
	for p != nil {
		if head.Val != p.Val {
			return false
		}
		p = p.Next
		head = head.Next
	}
	return true
}

// 也可以用快慢指针找到中间点
func isPalindrome2(head *ListNode) bool {
	// 找到链表中点的下一个位置
	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	// 反转后半段链表
	p = reverse(p)
	// 前半段与反转后的后半段对比
	q = head
	for p != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

// 实际上不应该破坏原链表，在判断完毕后，应该恢复原数组
func isPalindrome(head *ListNode) bool {
	// 找到链表中点的下一个位置
	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	m := p // 记录中点，最后恢复
	// 反转后半段链表
	p = reverse(p)
	// 前半段与反转后的后半段对比
	q = head
	for p != nil {
		if p.Val != q.Val {
			break
		}
		p = p.Next
		q = q.Next
	}
	// 后半段再次反转，恢复到原来
	_ = reverse(m)
	return p == nil
}

// 反转一个链表，并返回反转后的头节点
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

/* 值得一提的是以下递归解法，虽然空间复杂度是O(n)：
 */
func isPalindrome3(head *ListNode) bool {
	front := head

	var recusivelyCheck func(*ListNode) bool

	recusivelyCheck = func(current *ListNode) bool {
		if current == nil {
			return true
		}
		if !recusivelyCheck(current.Next) || current.Val != front.Val {
			return false
		}
		front = front.Next
		return true
	}
	return recusivelyCheck(head)
}
