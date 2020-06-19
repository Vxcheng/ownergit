package main

import (
	"testing"
)

func TestXML(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		var a int64 = 1
		t.Log("a: ", int(a))
	})
}
