package main

func demo1() {
	m := make(map[string]int)
	b := m["b"]
	_ = b
	_, ok := m["a"]
	_ = ok
}

func demo2() {
	m := make(map[string]int)
	m["b"] = 1
	for k, v := range m {
		_, _ = k, v
	}
}
