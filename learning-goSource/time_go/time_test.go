package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Run("", func(t *testing.T) {
		str := "2021-01-21 11:11:11"
		tt, _ := time.Parse("2006-01-02", str)
		tt.Format("20060102")
	})
}
