// main.go
package main

/*
// C 头文件
#include <stdio.h>
#include <stdlib.h>

// 声明 C 函数
int add(int a, int b);
void printMessage(char* message);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 导出给 C 调用的 Go 函数（需要在 C 代码中声明）
//
//export Multiply
func Multiply(a, b C.int) C.int {
	return a * b
}

/*
编译和运行
bash
# 编译 C 代码为静态库
gcc -c clib.c -o clib.o
ar rcs libclib.a clib.o

# 运行 Go 程序
go run .
*/

func main() {
	fmt.Println("=== Go 调用 C 函数 ===")

	// 调用 C 的 add 函数
	result := C.add(10, 20)
	fmt.Printf("C add(10, 20) = %d\n", result)

	// 调用 C 的 printMessage 函数
	message := C.CString("Hello from Go!")
	defer C.free(unsafe.Pointer(message)) // 记得释放内存
	C.printMessage(message)

	// 演示字符串处理
	goString := "Another message"
	cString := C.CString(goString)
	defer C.free(unsafe.Pointer(cString))
	C.printMessage(cString)
}
