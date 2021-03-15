package main

func toStack() int {
	a := new(int)
	*a = 1
	return *a
}

func toHeap() *int {
	x := 2
	return &x
}

type S struct {
	M *int
}

func main() {
	var a S
	x := &a
	_ = toStack_2(x)
	_ = toHeap_2(a)

	//
	var p S
	var f int
	ref(&f, &p)

	//
	var p1 S
	var f1 int
	ref(&f1, &p1)
}

func toStack_2(z *S) *S {
	return z
}

func toHeap_2(z S) *S {
	return &z
}

func ref(x *int, y *S) {
	y.M = x
}

func ref_2(x *int, y *S) {
	x = y.M
}
