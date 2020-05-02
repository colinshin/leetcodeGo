/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import (
	"container/heap"
	"sort"
)

/*
先统计每个任务的数量，贪心策略，需要先安排数量大的任务
以n+1个任务为一轮，同一轮中一个任务最多被安排一次。
每一轮中，将当前任务按照剩余次数降序排列，再选择剩余次数最多的n+1个任务一次执行
如果任务的种类 t 少于 n + 1 个，就只选择全部的 t 种任务，其余的时间空闲。

时间复杂度: O(result),给每个任务都安排了时间，因此时间复杂度和最终的答案成正比
空间复杂度: O(26)=O(1)
*/
func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	sort.Ints(count)
	result := 0
	for count[25] > 0 {
		for i := 0; i <= n && count[25] > 0; i++ {
			result++
			if i < 26 && count[25-i] > 0 {
				count[25-i]--
			}
		}
		sort.Ints(count)
	}
	return result
}

/*
继承自上面的方法，在选择每一轮任务时，可用堆代替排序。
一开始，把所有任务的数量加入堆。
每一轮，从堆里选择最多n+1个任务，把它们数量减去1，如果不为0，再重新放回堆中；直到堆为空

时空复杂度与上面一样
*/
func leastInterval1(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	h := &Heap{}
	for _, v := range count {
		if v > 0 {
			heap.Push(h, v)
		}
	}
	result := 0
	for h.Len() > 0 {
		var tmp []int
		for i := 0; i <= n && (h.Len() > 0 || len(tmp) > 0); i++ {
			result++
			if h.Len() == 0 {
				continue
			}
			t := heap.Pop(h).(int)
			if t > 1 {
				tmp = append(tmp, t-1)
			}
		}
		for _, v := range tmp {
			heap.Push(h, v)
		}
	}
	return result
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

/*
时间复杂度：O(M)，其中 M 是任务的总数，即tasks数组的长度。
空间复杂度：O(1)。
*/
func leastInterval2(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	sort.Ints(count)
	max := count[25] - 1
	idleSlots := max * n
	for i := 24; i >= 0 && count[i] > 0; i-- {
		idleSlots -= min(count[i], max)
	}
	result := len(tasks)
	if idleSlots > 0 {
		result += idleSlots
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
时间复杂度：O(M)，其中 M 是任务的总数，即tasks数组的长度。
空间复杂度：O(1)。
*/
func leastInterval3(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	// 统计每个任务的数量，数量最大的任务的数量及个数
	count := make([]int, 26)
	max, maxCount := 0, 0
	for _, v := range tasks {
		c := count[v-'A'] + 1
		count[v-'A'] = c
		if max < c {
			max = c
			maxCount = 1
		} else if max == c {
			maxCount++
		}
	}
	result := (max-1)*(n+1) + maxCount
	if result < len(tasks) {
		return len(tasks)
	}
	return result
}
