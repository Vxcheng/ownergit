// bridge.go
package main

/*
// C 头文件
#include <stdio.h>
#include <stdlib.h>

// 声明 C 函数
void cCallback(char* message, int value);


*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 导出给 C 调用的 Go 函数

//export ProcessData
func ProcessData(data C.int) C.int {
	fmt.Printf("Go: ProcessData received: %d\n", data)

	// 调用 C 回调函数
	msg := C.CString("Processing completed")
	defer C.free(unsafe.Pointer(msg))
	C.cCallback(msg, data*2)

	return data * 3
}

//export StringLength
func StringLength(str *C.char) C.int {
	goStr := C.GoString(str)
	fmt.Printf("Go: StringLength received: %s\n", goStr)
	return C.int(len(goStr))
}

//export StartProcessing
func StartProcessing() {
	fmt.Println("Go: StartProcessing called")

	// 模拟处理并调用 C 回调
	for i := 0; i < 3; i++ {
		msg := C.CString(fmt.Sprintf("Progress update %d", i))
		C.cCallback(msg, C.int(i))
		C.free(unsafe.Pointer(msg))
	}
}

//export goCallback
func goCallback(message *C.char, value C.int) {
	fmt.Printf("Go: Callback received: %s, %d\n", C.GoString(message), value)
}

func main() {
	// 这个 main 函数在编译为 C 库时不会被调用
	fmt.Println("This is a Go library for C")
}
