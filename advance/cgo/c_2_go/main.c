#include <stdio.h>

extern int sum(int a, int b);

int main() {
    printf("sum: %d\n", sum(1, 2));
    return 0;
}

