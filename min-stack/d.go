package min_stack

/*
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, 0)
		s.min = x
	} else {
		s.stack = append(s.stack, x-s.min)
		if x < s.min {
			s.min = x
		}
	}
}

func (s *MinStack) Pop() {
	n := len(s.stack)
	if n == 0 {
		return
	}
	last := s.stack[n-1]
	if last < 0 {
		s.min -= last
	}
	s.stack = s.stack[:n-1]
}

func (s *MinStack) Top() int {
	n := len(s.stack)
	if n == 0 {
		return 0
	}
	last := s.stack[n-1]
	if last < 0 {
		return s.min
	}
	return last + s.min
}

func (s *MinStack) GetMin() int {
	return s.min
}
