/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package ii

import (
	"math"
	"sort"
)

type interval struct {
	start, end int
}

type MyCalendarTwo struct {
	// 分别表示已经添加的所有日程和已有日程重复的时间段组成的列表
	calendar, overlap []interval
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (mc *MyCalendarTwo) Book(start int, end int) bool {
	if len(mc.overlap) > 0 {
		// 二分搜索出新日程在overlap里的位置,不存在的话找需要插入的位置
		pos := sort.Search(len(mc.overlap), func(i int) bool {
			return mc.overlap[i].start >= start
		})
		// 查看搜索出的位置附件有没有和当前日程重叠的部分
		if pos < len(mc.overlap) && mc.overlap[pos].start < end ||
			pos-1 >= 0 && mc.overlap[pos-1].end > start {
			return false
		}
	}
	for _, v := range mc.calendar {
		if max(v.start, start) < min(v.end, end) {
			it := interval{start: max(start, v.start), end: min(end, v.end)}
			pos := sort.Search(len(mc.overlap), func(i int) bool {
				return mc.overlap[i].start >= it.start
			})
			insert(&mc.overlap, it, pos)
		}
	}
	pos := sort.Search(len(mc.calendar), func(i int) bool {
		return mc.calendar[i].start >= start
	})
	it := interval{start: start, end: end}
	insert(&mc.calendar, it, pos)
	return true
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// 在s中将val插入索引i处，插入前i及其后边元素一一后移
func insert(s *[]interval, val interval, i int) {
	if i == len(*s) {
		*s = append(*s, val)
		return
	}
	*s = append(*s, interval{})
	_ = copy((*s)[i+1:len(*s)], (*s)[i:len(*s)-1])
	(*s)[i] = val
}
