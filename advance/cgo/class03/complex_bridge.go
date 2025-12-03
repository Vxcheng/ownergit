// complex_bridge.go
package main

/*
#include <stdio.h>
#include <stdlib.h>

// 定义 C 结构体
typedef struct {
    int id;
    char* name;
    double value;
} DataStruct;

// 声明 C 函数
DataStruct* createData(int id, char* name, double value);
void freeData(DataStruct* data);
void printData(DataStruct* data);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export ProcessStruct
func ProcessStruct(data *C.DataStruct) C.double {
	fmt.Printf("Go: Processing struct - ID: %d, Name: %s, Value: %.2f\n",
		data.id, C.GoString(data.name), data.value)

	// 修改数据
	newValue := data.value * 1.1
	return C.double(newValue)
}

//export CreateGoData
func CreateGoData(id C.int, name *C.char, value C.double) *C.DataStruct {
	// 在 Go 中创建 C 结构体
	data := (*C.DataStruct)(C.malloc(C.size_t(unsafe.Sizeof(C.DataStruct{}))))
	data.id = id
	data.name = C.CString(C.GoString(name)) // 复制字符串
	data.value = value * 2.0

	return data
}

//export FreeGoData
func FreeGoData(data *C.DataStruct) {
	if data != nil {
		if data.name != nil {
			C.free(unsafe.Pointer(data.name))
		}
		C.free(unsafe.Pointer(data))
	}
}

func main() {}
