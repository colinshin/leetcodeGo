package __median_of_two_sorted_arrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T)  {
	cases := []struct{
		a []int
		b []int
		expect float64
	}{
		{a:[]int{1, 3}, b:[]int{2}, expect:2.0},
		{a:[]int{1, 3}, b:[]int{2, 4}, expect:2.5},
		{a:[]int{1, 3}, b:[]int{2, 4, 8}, expect:3.0},
	}
	for _, c := range cases {
		r := findMedianSortedArrays(c.a, c.b)
		if c.expect - r > 0.0000001 {
			t.Error("expect:", c.expect, "got:", r)
		}
	}
}
