// clib.c
#include <stdio.h>
#include <stdlib.h>

// 声明 Go 函数（在 Go 代码中定义）
extern int Multiply(int a, int b);

// C 函数实现
int add(int a, int b) {
    printf("C: add function called with %d and %d\n", a, b);
    return a + b;
}

void printMessage(char* message) {
    printf("C: Received message: %s\n", message);
    
    // 演示调用 Go 函数
    int result = Multiply(5, 6);
    printf("C: Called Go Multiply(5, 6) = %d\n", result);
}