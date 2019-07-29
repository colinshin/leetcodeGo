package integer

import "testing"

func TestMaxMin(t *testing.T) {
	cases := []struct {
		a   int
		b   int
		max int
		min int
	}{
		{a: 0, b: 5, max: 5, min: 0},
		{a: -3, b: 5, max: 5, min: -3},
		{a: -3, b: -5, max: -3, min: -5},
		{a: 8, b: 5, max: 8, min: 5},
	}

	for _, c := range cases {
		min, max := Min(c.a, c.b), Max(c.a, c.b)
		if min != c.min {
			t.Errorf("a: %d, b: %d, expected min: %d, but got:%d", c.a, c.b, c.min, min)
		}
		if max != c.max {
			t.Errorf("a: %d, b: %d, expected max: %d, but got:%d", c.a, c.b, c.max, max)
		}
	}
}
