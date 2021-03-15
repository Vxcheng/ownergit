package main

import (
	"testing"
)

func TestReplaceIndividually(t *testing.T) {
	// 关闭{}zData{}服务, {}, ["cell02", "store"]

	tests := []struct {
		name     string
		old, sep string
		keywords []string
		expected string
	}{
		{
			name:     "example1",
			old:      "关闭{}zData{}服务",
			sep:      "{}",
			keywords: []string{"cell02", "store"},
			expected: "关闭cell02zDatastore服务",
		},
		{
			name:     "example2",
			old:      "关闭{}",
			sep:      "{}",
			keywords: []string{"cell02"},
			expected: "关闭cell02",
		},
		{
			name:     "example2",
			old:      "{}关闭",
			sep:      "{}",
			keywords: []string{"cell02"},
			expected: "cell02关闭",
		},
	}

	for _, tt := range tests {
		got := ReplacePlaceholder(tt.old, tt.sep, tt.keywords)
		if got != tt.expected {
			t.Fatalf("got is '%s', want is '%s'", got, tt.expected)
		}
	}

}

func TestConvert(t *testing.T) {
	want := float64(2048)
	got := convert("2047.9975")
	if got != want {
		t.Fail()
	}
}
