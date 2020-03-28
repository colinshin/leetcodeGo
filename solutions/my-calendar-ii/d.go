/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package my_calendar_ii

import (
	"math"
	"sort"
)

/*
731. 我的日程安排表 II
https://leetcode-cn.com/problems/my-calendar-ii/

实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，
注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。

当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生三重预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。
请按照以下步骤调用MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)

示例：
MyCalendar();
MyCalendar.book(10, 20); // returns true
MyCalendar.book(50, 60); // returns true
MyCalendar.book(10, 40); // returns true
MyCalendar.book(5, 15); // returns false
MyCalendar.book(5, 10); // returns true
MyCalendar.book(25, 55); // returns true
解释：
前两个日程安排可以添加至日历中。 第三个日程安排会导致双重预订，但可以添加至日历中。
第四个日程安排活动（5,15）不能添加至日历中，因为它会导致三重预订。
第五个日程安排（5,10）可以添加至日历中，因为它未使用已经双重预订的时间10。
第六个日程安排（25,55）可以添加至日历中，因为时间 [25,40] 将和第三个日程安排双重预订；
时间 [40,50] 将单独预订，时间 [50,55）将和第二个日程安排双重预订。

提示：
每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/my-calendar-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
朴素实现
内部维护两个集合，分别表示已经添加的所有日程和已有日程重复的时间段组成的列表
集合可以用list或slice —— 如果后续需求增加，要支持日程删除，那用map还是挺合适的~

时间复杂度O(n^2)， 空间复杂度O(n)
一个可能的优化是保持两个集合有序，
这样可以用二分法找到新日程应该插入的位置，并与其附近元素比较是否重叠
用切片可以满足二分查询，但是要插入新值的时候，需要把后续的元素一一向后移动
用每个日程区间的start作为排序和二分查找的依据
综合复杂度O(n^2)， 空间复杂度不变，依然为O(n)
*/
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
		if pos < len(mc.overlap) && mc.overlap[pos].start < end {
			return false
		}
		if pos-1 >= 0 && mc.overlap[pos-1].end > start {
			return false
		}
	}
	pos := sort.Search(len(mc.calendar), func(i int) bool {
		return mc.calendar[i].start >= start
	})
	if len(mc.calendar) > 0 {
		if pos < len(mc.calendar) && mc.calendar[pos].start < end {
			it := interval{start: max(start, mc.calendar[pos].start), end: min(end, mc.calendar[pos].end)}
			i := sort.Search(len(mc.overlap), func(i int) bool {
				return mc.overlap[i].start >= it.start
			})
			insert(&mc.overlap, it, i)
		}
		if pos-1 >= 0 && mc.calendar[pos-1].end > start {
			it := interval{start: max(start, mc.calendar[pos-1].start), end: min(end, mc.calendar[pos-1].end)}
			i := sort.Search(len(mc.overlap), func(i int) bool {
				return mc.overlap[i].start >= it.start
			})
			insert(&mc.overlap, it, i)
		}
	}
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
