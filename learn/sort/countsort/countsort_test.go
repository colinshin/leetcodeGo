/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package countsort

import (
	"reflect"
	"testing"
)

func Test_countSort(t *testing.T) {
	tests := []struct {
		s    []int
		want []int
	}{
		// 元素都是非负数
		{s: []int{5, 4, 3, 2, 1, 0}, want: []int{0, 1, 2, 3, 4, 5}},
		{s: []int{5, 3, 3, 5, 1, 0}, want: []int{0, 1, 3, 3, 5, 5}},
		{s: []int{5, 1}, want: []int{1, 5}},
		// 元素有正有负
		{s: []int{5, -1}, want: []int{-1, 5}},
		{s: []int{-5, -10}, want: []int{-10, -5}},
		{s: []int{5, -3, -3, 5, 1, 0}, want: []int{-3, -3, 0, 1, 5, 5}},
		// 元素不够分散的情况
		{s: []int{1, 1, 1000, 1, 1}, want: []int{1, 1, 1, 1, 1000}},
	}
	for _, tt := range tests {
		if got := countSort(tt.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("countSort() = %v, want %v", got, tt.want)
		}
	}
}
