package tdd

import "testing"

func TestHandleStr(t *testing.T) {
	cases := []struct {
		name    string
		context string
		want    string
	}{
		{name: "left extra open", context: "a((b)", want: "a(b)"},
		{name: "right extra close", context: "a(b))", want: "a(b)"},
		{name: "balanced", context: "a(b)", want: "a(b)"},
		{name: "no paren", context: "a", want: "a"},
		{name: "trailing open", context: "a(b)(", want: "a(b)"},
		{name: "leading open", context: "(ab()", want: "ab()"},
		{name: "unmatched close first", context: ")ab()", want: "ab()"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := handleStr1(c.context)
			if got != c.want {
				t.Fatalf("handleStr1(%q) = %q, want %q", c.context, got, c.want)
			}
		})
	}
}
