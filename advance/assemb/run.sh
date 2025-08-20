
go tool compile -N -l -S main.go
GOOS=linux GOARCH=amd64 go tool compile -N -l -S main.go # 指定系统和架构
go tool objdump #
方法1： 根据目标文件反编译出汇编代码

go tool compile -N -l main.go # 生成main.o
go tool objdump main.o
go tool objdump -s "main.(main|add)" ./test # objdump支持搜索特定字符串
方法2： 根据可执行文件反编译出汇编代码

go build -gcflags="-N -l" main.go -o test
go tool objdump main.o

go build -gcflags -S #
go build -gcflags="-N -l -S"  main.go

