package main

import "testing"

func TestMap(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		demo1()
	})

	t.Run("demo2", func(t *testing.T) {
		demo2()
	})

	t.Run("demo3", func(t *testing.T) {
		demo3()
	})
}
