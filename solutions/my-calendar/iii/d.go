/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package iii

import (
	"container/list"
	"math"
)

type point struct {
	pos  int // 该点在数轴上的位置。
	deep int // 该点的深度——即被多少条线段包含。
}

type MyCalendarThree struct {
	points *list.List
}

func Constructor() MyCalendarThree {
	mc := MyCalendarThree{points: list.New()}
	// 结合list特点, 方便后续处理，先预置两个点，无限小点和和无限大点
	mc.points.PushBack(&point{pos: -1, deep: 0})
	mc.points.PushBack(&point{pos: math.MaxInt64, deep: 0})
	return mc
}

func (mc *MyCalendarThree) Book(start int, end int) int {
	var startNode, endNode *list.Element
	// 插入起始点
	for e := mc.points.Front(); e != nil; e = e.Next() {
		if start == e.Value.(*point).pos { // 避免重复，如果该点已经存在了就不新建了
			startNode = e
			break
		}
		if start > e.Value.(*point).pos &&
			start < e.Next().Value.(*point).pos { // 新建一个点，注意新建点的颜色深度 暂时 和它前面的点的颜色深度一致
			startNode = mc.points.InsertAfter(&point{pos: start, deep: e.Value.(*point).deep}, e)
			break
		}
	}
	// 插入结束点。
	for e := mc.points.Back(); e != nil; e = e.Prev() {
		if end == e.Value.(*point).pos {
			endNode = e
			break
		}
		if end < e.Value.(*point).pos &&
			end > e.Prev().Value.(*point).pos {
			endNode = mc.points.InsertBefore(&point{pos: end, deep: e.Prev().Value.(*point).deep}, e)
			break
		}
	}
	// 对于起始和结束点之间的所有点，深度都加一。
	for e := startNode; e != endNode && e != nil; e = e.Next() {
		p := e.Value.(*point)
		p.deep++
	}

	k := 0
	for e := mc.points.Front(); e != nil; e = e.Next() {
		if e.Value.(*point).deep > k {
			k = e.Value.(*point).deep
		}
	}
	return k
}
