

2. 调度器调试
# 跟踪Goroutine调度事件
GODEBUG=schedtrace=1000,scheddetail=1 go run main.go

# 输出示例（每1000ms）：
# SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1...

3. GC 调试组合
pacer: assist ratio=+2.849426e+000 (scan 0 MB in 3->4 MB) workers=2++0.000000e+000
pacer: 25% CPU (25 exp.) for 442864+19112+466850 B work (466850 B exp.) in 4030464 B -> 4308992 B (∆goal 114688, cons/mark +0.000000e+000)
gc 1 @0.048s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 3->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
pacer: assist ratio=+3.384340e+000 (scan 1 MB in 3->4 MB) workers=2++0.000000e+000
pacer: 25% CPU (25 exp.) for 508392+19144+466850 B work (929394 B exp.) in 3919688 B -> 3944264 B (∆goal -250040, cons/mark +9.995701e-002)
gc 2 @0.073s 0%: 0+0.53+0 ms clock, 0+0/0/0+0 ms cpu, 3->3->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P

1. 并发问题诊断
bash
# 检测共享变量访问（竞态检测增强版）
GODEBUG=inittrace=1,clobberfree=1 go run -race main.go


1. Goroutine泄漏诊断
bash
GODEBUG=scheddetail=1,gctrace=1 go run main.go
# 配合分析：
go tool pprof http://localhost:6060/debug/pprof/goroutine
2. 内存增长分析
bash
GODEBUG=allocfreetrace=1,gctrace=1 go run main.go
# 配合分析：
go tool pprof -http=:8080 -alloc_space mem.pprof
3. 调度延迟分析
bash
GODEBUG=schedtrace=100,scheddetail=1 go run main.go
# 配合分析：
go tool trace trace.out



