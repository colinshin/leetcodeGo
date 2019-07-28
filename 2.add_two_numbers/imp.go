package __add_two_numbers

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func NewWithArray(array []int) *ListNode {
	result := &ListNode{}
	current := result
	for i, v := range array {
		current.Val = v
		if i < len(array) -1 {
			current.Next = &ListNode{}
		}
		current = current.Next
	}
	return result
}

func (list *ListNode) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("(")
	for list != nil {
		buffer.WriteString(strconv.Itoa(list.Val))
		if list.Next != nil {
			buffer.WriteString(" -> ")
		}
		list = list.Next
	}
	buffer.WriteString(")")
	return buffer.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	carry := 0 // must be 0 or 1
	for l1 !=nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Val = sum % 10
		if l1 !=nil || l2 != nil || carry != 0 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return result
}
