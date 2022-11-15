package main

type S struct {
	M *int
}

func main() {
	var i int
	refStruct(i)
}

func refStruct(y int) (z S) {
	z.M = &y
	return z
}
