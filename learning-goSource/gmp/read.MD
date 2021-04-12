go tool trace trace.out
GODEBUG=schedtrace=1000 go run debug.go
go tool compile -S debug.go > debug_compile.S