/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package i

import (
	"container/list"
	"math"
)

const (
	kindStart = 1
	kindEnd   = -1
)

type point struct {
	pos, kind int
}

type MyCalendar struct {
	points *list.List
}

func Constructor() MyCalendar {
	mc := MyCalendar{points: list.New()}
	mc.points.PushBack(&point{pos: -1})
	mc.points.PushBack(&point{pos: math.MaxInt64})
	return mc
}

func (mc *MyCalendar) Book(start int, end int) bool {

	var startNode, endNode *list.Element
	for e := mc.points.Front(); e != nil; e = e.Next() {
		if start == e.Value.(*point).pos && e.Value.(*point).kind == kindStart {
			return false
		}
		if start >= e.Value.(*point).pos && start < e.Next().Value.(*point).pos {
			startNode = mc.points.InsertAfter(&point{pos: start, kind: kindStart}, e)
			break
		}
	}
	for e := mc.points.Back(); e != nil; e = e.Prev() {
		if end == e.Value.(*point).pos && e.Value.(*point).kind == kindEnd {
			return false
		}
		if end <= e.Value.(*point).pos && end > e.Prev().Value.(*point).pos {
			endNode = mc.points.InsertBefore(&point{pos: end, kind: kindEnd}, e)
			break
		}
	}
	for e := startNode.Next(); e != endNode; e = e.Next() {
		p := e.Value.(*point)
		if p.kind == kindStart || p.kind == kindEnd {
			_ = mc.points.Remove(startNode)
			_ = mc.points.Remove(endNode)
			return false
		}
	}
	return true
}
