package main

import (
	"reflect"
	"regexp"
	"testing"
)

type Class struct {
	F string
}

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

func TestPoint(t *testing.T) {
	t.Run("", func(t *testing.T) {
		a, b := 1, 2
		swap1(&a, &b)
		t.Log(a, b)
	})

	t.Run("", func(t *testing.T) {
		a, b := 1, 2
		swap2(&a, &b)
		t.Log(a, b)
	})

	tests := []struct {
		name string
	}{}

	for _ = range tests {
		// t.Run(tt.name, func(t *testing.T) {

		// })
	}

}

func swap1(a, b *int) {
	*a = *b
}

func swap2(a, b *int) {
	tmp := *b
	*a = tmp
}

func TestMatch(t *testing.T) {
	str := "   Active: active (running) since Thu 2021-09-16 19:18:20 CST; 24s ago"
	re := regexp.MustCompile(`Active:\s+active\s+\(running\)`)
	if !re.MatchString(str) {
		t.Error()
	}
}

func TestUnionParams(t *testing.T) {
	params := unionParams(map[string]string{"a": "a"}, map[string]string{"b": "b"})
	t.Log(params)

	ss := splitSlice([]string{"a", "b", "c", "d", "e"})
	t.Log(ss)
}

func splitSlice(in []string) []string {
	for i := range in {
		if i == 1 || i == 3 {
			in = append(in[:i], in[i+1:]...)
			continue
		}
	}

	return in
}

func unionParams(staticParams, dynamicParams map[string]string) map[string]string {
	for k, v := range dynamicParams {
		staticParams[k] = v
	}

	return staticParams
}
