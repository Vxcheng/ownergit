package main

import (
	"fmt"
	"testing"
)

func Test_reverseElement(t *testing.T) {
	tests := []struct {
		name   string
		Linked *ListNode
	}{
		// TODO: Add test cases.
		struct {
			name   string
			Linked *ListNode
		}{
			"reverse list len is 5",
			&ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							&ListNode{5, nil}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseElement(tt.Linked, 2, 4)
			fmt.Printf("got is: %v\n", *got)
		})
	}
}
