package main

/*
#include <stdint.h>

int64_t myadd(int64_t a, int64_t b) {
    return a+b;
}
*/
import "C"

import (
	asmpkg "ownergit/advance/assemb/class04/pkg"
	"runtime"
	"unsafe"
)

func main() {
	if runtime.GOOS != "windows" {
		println(asmpkg.AsmCallCAdd(
			uintptr(unsafe.Pointer(C.myadd)),
			123, 456,
		))
	}
}
