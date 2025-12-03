// complex_main.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*
echo "编译 Go 代码为 C 静态库..."
go build -buildmode=c-archive -o libgobridge.a complex_bridge.go

echo "编译 C 主程序..."
gcc -o c_program complex_.c libgobridge.a -lpthread

echo "运行 C 程序..."
./c_program
*/

// 声明 Go 函数
extern double ProcessStruct(void* data);
extern void* CreateGoData(int id, char* name, double value);
extern void FreeGoData(void* data);

// C 结构体定义
typedef struct {
    int id;
    char* name;
    double value;
} DataStruct;

// C 函数实现
DataStruct* createData(int id, char* name, double value) {
    DataStruct* data = (DataStruct*)malloc(sizeof(DataStruct));
    data->id = id;
    data->name = strdup(name);
    data->value = value;
    return data;
}

void freeData(DataStruct* data) {
    if (data) {
        free(data->name);
        free(data);
    }
}

void printData(DataStruct* data) {
    printf("C: Data - ID: %d, Name: %s, Value: %.2f\n", 
           data->id, data->name, data->value);
}

int main() {
    printf("=== 复杂数据类型示例 ===\n");
    
    // 创建 C 数据并传递给 Go
    DataStruct* cData = createData(1, "Test Data", 100.5);
    printData(cData);
    
    double result = ProcessStruct(cData);
    printf("C: Go returned: %.2f\n", result);
    
    // 测试 Go 创建的数据
    DataStruct* goData = (DataStruct*)CreateGoData(2, "Go Data", 50.0);
    printData(goData);
    FreeGoData(goData);
    
    freeData(cData);
    return 0;
}
