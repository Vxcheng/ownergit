package main

import (
	"regexp"
	"testing"
)

func TestExp1(t *testing.T) {
	tests := []struct {
		value string
	}{
		{
			value: "Drive /c0/e252/s5",
		},
		{
			value: "Drive /c0/e252/s5 - Detailed Information",
		},
	}

	for _, tt := range tests {
		t.Run("FindStringSubmatch", func(t *testing.T) {
			re := regexp.MustCompile(`(?i:drive) /c(\d+)/e(\d+)/s(\d+)$`)
			outs := re.FindStringSubmatch(tt.value)
			t.Log(outs)
		})
	}

}

func TestExp2(t *testing.T) {
	tests := []struct {
		value string
	}{
		{
			value: "Drive /c0/e252/s5",
		},
		{
			value: "Drive /c0/e252/s5 - Detailed Information",
		},
	}

	for _, tt := range tests {
		t.Run("FindStringSubmatch", func(t *testing.T) {
			re := regexp.MustCompile(`^(?i)(drive\s+/c\d+/e\d+/s\d+)\s+-\s+(?i:detailed information)$`)
			outs := re.FindStringSubmatch(tt.value)
			t.Log(outs)
		})
	}
}
