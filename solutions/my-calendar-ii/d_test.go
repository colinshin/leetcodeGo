/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package my_calendar_ii

import (
	"sort"
	"testing"
)

func TestMyCalendarTwo_Book(t *testing.T) {
	s := []interval{{start: 5, end: 19}, {start: 27, end: 40}}
	i := sort.Search(len(s), func(i int) bool {
		return s[i].start >= 3
	})
	t.Log(i)
	if i < len(s) && s[i].start < 14 {
		t.Log("true")
	}

	//cases := []struct {
	//	intervals [][]int
	//	want      []bool
	//}{
	//	{
	//		intervals: [][]int{{10, 20}, {50, 60}, {10, 40}, {5, 15}},
	//		want:      []bool{true, true, true, false},
	//	},
	//	{
	//		intervals: [][]int{{24, 40}, {43, 50}, {27, 43}, {5, 21}, {30, 40}},
	//		want:      []bool{true, true, true, true, false},
	//	},
	//}
	//for _, c := range cases {
	//	myCal := Constructor()
	//	for i, v := range c.intervals {
	//		got := myCal.Book(v[0], v[1])
	//		if got != c.want[i] {
	//			t.Errorf("%v, got %v, want %v", v, got, c.want[i])
	//		}
	//	}
	//}
}

func Test_insert(t *testing.T) {
	tests := []struct {
		s   *[]interval
		val interval
		i   int
	}{
		{
			s:   &[]interval{{start: 27, end: 40}},
			val: interval{start: 5, end: 19},
			i:   0,
		},
	}
	for _, tt := range tests {
		insert(tt.s, tt.val, tt.i)
		t.Log(tt.s)
	}
}
