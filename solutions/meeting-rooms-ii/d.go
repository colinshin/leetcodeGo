/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package meeting_rooms_ii

import (
	"sort"
)

/*
<253>
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，
为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

示例 1:
输入: [[0, 30],[5, 10],[15, 20]]
输出: 2

示例 2:
输入: [[7,10],[2,4]]
输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/meeting-rooms-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[i][1]
	})
	h := &IntHeap{}

	result := 1
	var tmp [][]int
	for _, v := range intervals {
		result = max(result, update(&tmp, v))
	}
	return result
}

func update(intervals *[][]int, interval []int) int {
	r := 0
	for _, v := range *intervals {
		if max(v[0], interval[0]) < min(v[1], interval[1]) {
			r++
		}
	}
	*intervals = append(*intervals, interval)
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
