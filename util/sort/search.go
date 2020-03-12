/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sort

// Search uses binary search to find and return the smallest index i
// in [from, to) at which f(i) is true, assuming that on the range [from, to),
// f(i) == true implies f(i+1) == true. That is, Search requires that
// f is false for some (possibly empty) prefix of the input range [from, to)
// and then true for the (possibly empty) remainder; Search returns
// the first true index. If there is no such index, Search returns n.
// It is just another version of the same name function in standard lib, package sort.
// this version is a little different of parameters
func Search(from, to int, f func(int) bool) int {
	for from < to {
		mid := int(uint(from+to) >> 1) // avoid overflow
		// from ≤ mid < to
		if !f(mid) {
			from = mid + 1 // preserves f(from-1) == false
		} else {
			to = mid // preserves f(to) == true
		}
	}
	// from == to, f(to-1) == false, and f(to) == true  =>  answer is from (or to).
	return from
}
