/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package stack_with_queen

import "container/list"

type MyStack struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list: list.New()}
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.list.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	if s.list.Len() == 0 {
		return -1
	}
	last := s.list.Back()
	r := last.Value.(int)
	s.list.Remove(last)
	return r
}

/** Get the top element. */
func (s *MyStack) Top() int {
	if s.list.Len() == 0 {
		return -1
	}
	return s.list.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return s.list.Len() == 0
}
