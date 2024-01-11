package main_test

import (
	"math/rand"
	"testing"
	"time"
)

func TestAAA(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		for i:=0; i<10;i++{
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			v := r.Intn(10)
			t.Log(v)
		}
	})
}