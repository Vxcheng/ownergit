#include <stdio.h>

extern int sum(int a, int b);

/*

echo "编译 Go 代码为 C 静态库..."
go build -buildmode=c-archive -o bridge.a main.go

echo "编译 C 主程序..."
gcc -o c_program main.c bridge.a -lpthread

echo "运行 C 程序..."
./c_program
*/
int main() {
    printf("sum: %d\n", sum(1, 2));
    return 0;
}

