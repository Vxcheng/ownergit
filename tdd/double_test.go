package tdd

import (
	"reflect"
	"testing"
)

func TestDoubleSum(t *testing.T) {
	cases := []struct {
		name string
		a, b [][2]int
		want [][2]int
	}{
		{
			name: "example with padding zeros",
			a:    [][2]int{{3, 3}, {4, 11}},
			b:    [][2]int{{2, 1}, {4, 2}, {3, 2}},
			want: [][2]int{{6, 1}, {12, 4}, {0, 9}},
		},
		{
			name: "left shorter padded zeros",
			a:    [][2]int{{5, 2}},
			b:    [][2]int{{2, 1}, {7, 3}},
			want: [][2]int{{10, 1}, {35, 1}, {0, 2}},
		},
		{
			name: "both empty",
			a:    nil,
			b:    nil,
			want: nil,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := doubleSum(c.a, c.b)
			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("doubleSum(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
			}
		})
	}
}
