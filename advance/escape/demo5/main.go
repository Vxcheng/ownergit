package main

type S struct {
	M *int
}

func main() {
	var x S
	var i int
	ref(&i, &x)
}

func ref(y *int, z *S) {
	z.M = y
}
