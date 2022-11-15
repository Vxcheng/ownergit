package main

// go tool compile -S main.go
func main() {
	m := make(map[string]int)
	b := m["b"]
	_ = b
}
