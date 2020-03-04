/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package others

/*
最直观的是用一个bool数组模拟这一排座位；坐了人标为true，没坐人标为false
*/
type ExamRoom struct {
	seated []bool
	count  int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seated: make([]bool, N),
		count:  0,
	}
}

func (room *ExamRoom) Seat() int {
	if room.count == 0 { // 还没有人入座，直接将0插入
		room.seated[0] = true
		room.count++
		return 0
	}
	maxDist := 0
	prevSeated := -1
	target := -1
	for i := 0; i < len(room.seated); i++ {
		if room.seated[i] {
			prevSeated = i
			break
		}
	}
	for i := prevSeated + 1; i < len(room.seated); i++ {
		if !room.seated[i] {
			continue
		}
		halfDist := (i - prevSeated) / 2
		if halfDist > maxDist {
			maxDist = halfDist
			target = prevSeated + halfDist
		}
		prevSeated = i
	}
	if target == -1 { // 只有一个座位坐了人
		if prevSeated+1 > len(room.seated)/2 {
			target = 0
		} else {
			target = len(room.seated) - 1
		}
	}
	room.seated[target] = true
	room.count++
	return target
}

func (room *ExamRoom) Leave(p int) {
	if !room.seated[p] {
		return
	}
	room.seated[p] = false
	room.count--
}
