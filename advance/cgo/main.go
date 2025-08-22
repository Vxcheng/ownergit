package main

/*
#include <stdint.h>
#include <errno.h>

//div
int div(int a, int b) {
	if (b == 0) {
		errno = EINVAL;
		return 0;
	}

	return a / b;
}

// 定义一个 C 语言函数，用于两个整数相加
int add(int a, int b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
	// 定义两个 Go 语言中的整数
	a, b := 3, 4

	// 调用 C 语言中的 add 函数
	result := C.add(C.int(a), C.int(b))

	// 将结果从 C 类型转换为 Go 类型，并打印出来
	fmt.Printf("The result of adding %d and %d is %d\n", a, b, int(result))

	v, err := C.div(2, 1)
	fmt.Println(v, err)

	v, err = C.div(1, 0)
	fmt.Println(v, err)
}
