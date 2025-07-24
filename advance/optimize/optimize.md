编译选项
Go 编译器默认编译出来的程序会带有符号表和调试信息，一般来说 release 版本可以去除调试信息以减小二进制体积。

$ go build -ldflags="-s -w" -o server main.go
$ ls -lh server
-rwxr-xr-x  1 dj  staff   7.8M Dec  8 00:29 server
-s：忽略符号表和调试信息。
-w：忽略DWARFv3调试信息，使用该选项后将无法使用gdb进行调试。
体积从 9.8M 下降到 7.8M，下降约 20%。

upx 是一个常用的压缩动态库和可执行文件的工具，通常可减少 50-70% 的体积。

$ go build -ldflags="-s -w" -o server main.go && upx -9 server

// +build debug 表示 build tags 中包含 debug 时，该源文件参与编译。
// +build !debug 表示 build tags 中不包含 debug 时，该源文件参与编译。
