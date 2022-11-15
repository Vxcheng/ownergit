package main

type S struct{}

func main() {
	var x S
	y := &x
	_ = *identity(y)
}

func identity(z *S) *S {
	return z
}
