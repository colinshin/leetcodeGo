package maximum

import "container/list"

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。



进阶：

你能在线性时间复杂度内解决此题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxSlidingWindow(nums []int, k int) []int {
	window := NewWindow()
	var result []int
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.Enqueue(nums[i])
			continue
		}
		window.Enqueue(nums[i])
		result = append(result, window.Max())
		window.Dequeue(i - k + 1)
	}
	return result
}

type Window struct {
	*list.List
}

func NewWindow() *Window {
	return &Window{list.New()}
}

func (w *Window) Enqueue(v int) {
	for w.Len() > 0 && w.Back().Value.(int) < v {
		_ = w.Remove(w.Back())
	}
	w.PushBack(v)
}

func (w *Window) Dequeue(v int) {
	if w.Len() > 0 && w.Front().Value.(int) <= v {
		_ = w.Remove(w.Front())
	}
}

func (w *Window) Max() int {
	return w.Front().Value.(int)
}

func maxSlidingWindow1(nums []int, k int) []int {
	var ret, windows []int
	for i, v := range nums {
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}
		for len(windows) > 0 && nums[windows[len(windows)-1:][0]] <= v {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, i)
		if i >= k-1 {
			ret = append(ret, nums[windows[0]])
		}
	}
	return ret
}
