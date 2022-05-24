package leetcode

import (
	"testing"
)

func TestMerge(t *testing.T) {

	// [1,3],[2,6],[8,15],[10,18]
	list := [][]int{
		{1, 3}, {2, 6}, {8, 15}, {10, 18},
	}
	got := merge(list)
	t.Log(got)
}

func TestLongestDupSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			input: "banana",
		},
		{
			input: "aaaabc",
		},
		{
			input: "hello",
		},
		{
			input: "abcdabcd卡卡",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestDupSubstring(tt.input)
			println("value:", got)
			t.Log(got)
		})
	}
}

func TestMergeDemo(t *testing.T) {
	t.Run("", func(t *testing.T) {
		list := [][]int{
			{1, 3}, {2, 6}, {8, 15}, {10, 18},
		}
		got := merge1(list)
		t.Log(got)
	})
}

func merge1(intervals [][]int) (ret [][]int) {
	// 按列遍历
	for i := 0; i < len(intervals[i])-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			// 判断重叠
			if intervals[i][0] < intervals[i+1][0] {
				// intervals[i][0] = intervals[i][0]
			} else {
				intervals[i][0] = intervals[i+1][0]
			}

			if intervals[i][1] > intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			} else {
				intervals[i][1] = intervals[i+1][1]
			}

		}

	}
	return intervals
}

func TestPara(t *testing.T) {
	para()
}
