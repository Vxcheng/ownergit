package main

type S struct{}

func main() {
	var x S
	_ = *ref(x)
}

func ref(z S) *S {
	return &z
}
