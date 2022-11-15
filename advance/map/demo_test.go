package main

import "testing"

func TestMap(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		demo1()
	})

	t.Run("demo2", func(t *testing.T) {
		demo2()
	})
}
