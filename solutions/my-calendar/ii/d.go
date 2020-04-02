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
			// 在pos处插入it
			mc.overlap = append(append(mc.overlap[:pos:pos], it), mc.overlap[pos:]...)
		}
	}
	pos := sort.Search(len(mc.calendar), func(i int) bool {
		return mc.calendar[i].start >= start
	})
	it := interval{start: start, end: end}
	// 在pos处插入it
	mc.calendar = append(append(mc.calendar[:pos:pos], it), mc.calendar[pos:]...)
	return true
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
