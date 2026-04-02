package leetcode

import (
	"sort"
	"strings"
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

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "basic two sums",
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			name: "single element match",
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},
		{
			name: "contains negative and zero",
			nums: []int{1, -1, 0},
			k:    0,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("subarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "classic example",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "single char exact",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "duplicate target chars",
			s:    "aa",
			t:    "aa",
			want: "aa",
		},
		{
			name: "end match",
			s:    "ab",
			t:    "b",
			want: "b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "basic anagram groups",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{
			name: "single string",
			strs: []string{"abc"},
			want: [][]string{{"abc"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if !equalGroupAnagrams(got, tt.want) {
				t.Fatalf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}

func TestGroupAnagrams1(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "basic anagram groups",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{
			name: "single string",
			strs: []string{"abc"},
			want: [][]string{{"abc"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams1(tt.strs)
			if !equalGroupAnagrams(got, tt.want) {
				t.Fatalf("groupAnagrams1(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}

func equalGroupAnagrams(a, b [][]string) bool {
	normalize := func(groups [][]string) []string {
		if groups == nil {
			return nil
		}
		result := make([]string, 0, len(groups))
		for _, group := range groups {
			sorted := append([]string(nil), group...)
			sort.Strings(sorted)
			result = append(result, strings.Join(sorted, ","))
		}
		sort.Strings(result)
		return result
	}

	aNorm := normalize(a)
	bNorm := normalize(b)
	if len(aNorm) != len(bNorm) {
		return false
	}
	for i := range aNorm {
		if aNorm[i] != bNorm[i] {
			return false
		}
	}
	return true
}
