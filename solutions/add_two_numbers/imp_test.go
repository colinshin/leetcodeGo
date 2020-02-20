/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package add_two_numbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	n1 := NewWithArray([]int{2, 4, 3})
	n2 := NewWithArray([]int{5, 6, 4})
	r := addTwoNumbers(n1, n2)
	t.Log(n1)
	t.Log(n2)
	t.Log(r)
	expected := NewWithArray([]int{7, 0, 8})
	if r.String() != expected.String() {
		t.Error("expect:", expected, "got:", r)
	}
}
