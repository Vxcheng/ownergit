package main

/*
指针
栈大小超出
闭包
interface{}类型
*/
type S struct{}

func main() {
	var x S
	y := &x
	_ = *identity(y)
}

func identity(z *S) *S {
	return z
}
