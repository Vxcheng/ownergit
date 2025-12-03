// c_main.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 声明 Go 函数
extern void ProcessData(int data);
extern int StringLength(char* str);
extern void StartProcessing();
// 声明要在 Go 中实现的回调函数
extern void goCallback(char* message, int value);

// C 回调函数
void cCallback(char* message, int value) {
    printf("C Callback: %s, value: %d\n", message, value);
}

/*
echo "编译 Go 代码为 C 静态库..."
go build -buildmode=c-archive -o libgobridge.a bridge.go

echo "编译 C 主程序..."
gcc -o c_program c_main.c libgobridge.a -lpthread

echo "运行 C 程序..."
./c_program
*/
// C 主程序
int main() {
    printf("=== C 主程序调用 Go 函数 ===\n");
    goCallback( "start", 1);

    
    // 调用 Go 的 StringLength 函数
    int length = StringLength("Hello World");
    printf("C: String length: %d\n", length);
    
    // 调用 Go 的 ProcessData 函数
    printf("C: Calling ProcessData...\n");
    ProcessData(42);
    
    // 调用 Go 的 StartProcessing 函数
    printf("C: Calling StartProcessing...\n");
    StartProcessing();
    
    return 0;
}