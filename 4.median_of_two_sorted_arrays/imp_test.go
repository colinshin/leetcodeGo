package __median_of_two_sorted_arrays

import (
	"testing"
)

var cases = []struct {
	a      []int
	b      []int
	expect float64
}{
	{a: []int{1, 3}, b: []int{2}, expect: 2.0},
	{a: []int{2}, b: []int{1, 3}, expect: 2.0},
	{a: []int{1, 3}, b: []int{1, 3}, expect: 2.0},
	{a: []int{1, 3}, b: []int{2, 4}, expect: 2.5},
	{a: []int{1, 2}, b: []int{3, 4}, expect: 2.5},
	{a: []int{}, b: []int{1, 3}, expect: 2.0},
	{a: nil, b: []int{1, 3}, expect: 2.0},
	{a: []int{1, 3}, b: []int{2, 4, 8}, expect: 3.0},
	{a: []int{1, 3}, b: nil, expect: 2.0},
	{a: []int{1, 3, 5}, b: []int{8, 10}, expect: 5.0},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, c := range cases {
		r := findMedianSortedArrays(c.a, c.b)
		if c.expect != r {
			t.Error("expect:", c.expect, "got:", r)
		}
	}
}

func TestFindMedianSortedArrays1(t *testing.T) {
	for _, c := range cases {
		r := findMedianSortedArrays1(c.a, c.b)
		if c.expect != r {
			t.Error("expect:", c.expect, "got:", r)
		}
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	for _, c := range cases {
		r := findMedianSortedArrays2(c.a, c.b)
		if c.expect != r {
			t.Error("expect:", c.expect, "got:", r)
		}
	}
}
