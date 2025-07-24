package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestCompare(t *testing.T) {
	a := `["rac048(192.168.10.48)","rac049(192.168.10.49)"]`
	b := `["rac048(192.168.10.48)"`
	if a < b {
		t.Error()
	}

	ss := []string{a, b}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})

	t.Log(ss)
}

func TestArch(t *testing.T) {
	armArchRe := regexp.MustCompile(`aarch64`)

	tests := []struct {
		name string
		arch string
	}{
		{
			name: "x86",
			arch: "3.10.0-1160.el7.x86_64",
		},
		{
			name: "arm",
			arch: "4.19.90-17.ky10.aarch64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !armArchRe.MatchString(tt.arch) {
				t.Error()
			} else {
				t.Logf("%s match", tt.name)
			}
		})
	}
}

func TestInter(t *testing.T) {
	tests := []struct {
		value float64
		want  int64
	}{
		{
			value: 1.2,
			want:  1,
		},
		{
			value: 1.9,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := int64(tt.value)
			if got != tt.want {
				t.Error()
			}
		})
	}
}

func TestSplit(t *testing.T) {
	t.Run("", func(t *testing.T) {
		RAID := "RAID"

		out := strings.Split(fmt.Sprintf("%s1", "r"), RAID)
		t.Log(out)
		raidRe := regexp.MustCompile(`RAID\s*([\w]+)`)
		values := raidRe.FindStringSubmatch(fmt.Sprintf("%s1", RAID))
		t.Log(values)
	})
}
func TestStrconv(t *testing.T) {
	want := int64(2048)
	got, err := strconv.ParseInt(fmt.Sprintf("%f", 2048.0), 10, 64)
	assert.Nil(t, err)
	if got != want {
		t.Fail()
	}
}

func TestMapPoint(t *testing.T) {
	t.Run("notUseMapPoint", func(t *testing.T) {

		dict := make(map[string]int)
		for i := 0; i < 4; i++ {
			notUseMapPoint(dict)
		}

		t.Log(dict)
	})

	t.Run("useMapPoint", func(t *testing.T) {
		dict := make(map[string]*int)
		useMapPoint(dict)
		t.Log(dict)
	})
}

func useMapPoint(dict map[string]*int) {
	tmp := 2
	dict["a"] = &tmp
}

func notUseMapPoint(dict map[string]int) {
	dict["a"] = 1
}

func appendIgnoreCheckText(desc string) string {
	r := []rune(desc)
	if len(r) > 0 {
		if string(r[len(r)-1]) == "。" {
			tmp := r[:len(r)-1]
			desc = string(append(tmp, []rune("，")...))
		}
		return fmt.Sprintf("%s忽略该检查。", desc)
	} else {
		return desc
	}
}

func TestAppendIgnoreCheckText(t *testing.T) {
	tests := []struct {
		name string
		desc string
	}{
		{desc: "Manager中配置。"},
		{desc: "Manager中"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendIgnoreCheckText(tt.desc)
			t.Log(got)
		})
	}
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
