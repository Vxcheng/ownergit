package main

import (
	"reflect"
	"testing"
)

type Class struct{}

func TestType(t *testing.T) {
	var a Class
	b := Class{}
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("a: %v, b: %s", a, b)
	}

	var c *Class
	d := &Class{}
	if !reflect.DeepEqual(c, d) {
		t.Fatalf("c: %v, d: %s", c, d)
	}
}

func TestAppend(t *testing.T) {
	Append()
}
