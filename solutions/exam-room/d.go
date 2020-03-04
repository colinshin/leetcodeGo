/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package exam_room

import "container/list"

/*
在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。

当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)

返回 ExamRoom(int N) 类，它有两个公开的函数：其中，函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。



示例：

输入：["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
输出：[null,0,9,4,2,null,5]
解释：
ExamRoom(10) -> null
seat() -> 0，没有人在考场里，那么学生坐在 0 号座位上。
seat() -> 9，学生最后坐在 9 号座位上。
seat() -> 4，学生最后坐在 4 号座位上。
seat() -> 2，学生最后坐在 2 号座位上。
leave(4) -> null
seat() -> 5，学生最后坐在 5 号座位上。


提示：

1 <= N <= 10^9
在所有的测试样例中 ExamRoom.seat() 和 ExamRoom.leave() 最多被调用 10^4 次。
保证在调用 ExamRoom.leave(p) 时有学生正坐在座位 p 上。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/exam-room
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
最直观的是用一个bool数组模拟这一排座位；坐了人标为true，没坐人标为false
*/
type ExamRoom struct {
	seated *list.List // 表示坐着同学的位置
	last   int        // 最后一个座位， 就是总座位数减一
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seated: list.New(),
		last:   N - 1,
	}
}

func (room *ExamRoom) Seat() int {
	if room.seated.Len() == 0 { // 还没有人入座，直接将0插入
		room.seated.PushFront(0)
		return 0
	}
	e := room.seated.Front()
	pre := e.Value.(int)
	max := pre // 头部需要特殊判断
	addVal := 0
	addFront := e
	e = e.Next()
	for ; e != nil; e = e.Next() {
		val := e.Value.(int)
		distant := (val - pre) / 2 // 两点之间的最远距离
		if distant > max {
			max = distant
			addFront = e           // 需要插入的点的后一个元素。方便找到后直接插入
			addVal = pre + distant // 需要插入的位置
		}
		pre = val
	}
	if room.last-pre > max { // 尾部特殊判断
		room.seated.PushBack(room.last) // 直接插入到链表尾部
		return room.last
	}
	room.seated.InsertBefore(addVal, addFront) // 插入
	return addVal
}

func (room *ExamRoom) Leave(p int) {
	for e := room.seated.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			room.seated.Remove(e)
			return
		}
	}
	return
}
