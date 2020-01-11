package cherry_pickup

import "testing"

func Test_cherryPickup(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{name: "case 1", grid: [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickup(tt.grid); got != tt.want {
				t.Errorf("cherryPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
