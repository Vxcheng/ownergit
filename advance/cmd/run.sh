$ go build -n main.go
mkdir -p $WORK\b001\
cat >$WORK\b001\importcfg.link 
$ go build -work -a -p 1 -x main.go

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


自举编译
从官网下载发行包 #
第一种方式是从Go发行包中获取Go二进制应用，比如要源码编译go1.14.13，我们可以去 官网下载已经编译好的go1.13，设置好GOROOT_BOOTSTRAP环境变量，就可以源码编译了。

wget https://golang.org/dl/go1.13.15.linux-amd64.tar.gz
tar xzvf go1.13.15.linux-amd64.tar.gz
mv go go1.13.15
export GOROOT_BOOTSTRAP=/tmp/go1.13.15 # 设置GOROOT_BOOTSTRAP环境变量指向bootstrap toolchain的目录

cd /tmp
git clone -b go1.14.13 https://go.googlesource.com/go go1.14.13
cd go1.14.13/src
./make.bash
使用gccgo工具编译 #
第二种方式是使用gccgo来编译：

sudo apt-get install gccgo-5
sudo update-alternatives --set go /usr/bin/go-5
export GOROOT_BOOTSTRAP=/usr

cd /tmp
git clone -b go1.14.13 https://go.googlesource.com/go go1.14.13
cd go1.14.13/src
./make.bash
基于go1.14版本工具链编译 #
第三种方式是先编译出go1.4版本，然后使用go1.4版本去编译其他版本。

cd /tmp
git clone -b go1.4.3 https://go.googlesource.com/go go1.4
cd go1.4/src
./all.bash # go1.4版本是c语言实现的编译器
export GOROOT_BOOTSTRAP=/tmp/go1.4


git clone -b go1.14.13 https://go.googlesource.com/go go1.14.13
cd go1.14.13/src
./all.bash


go tool compile -N -l main.go # 生成main.o
go tool objdump main.o # 打印所有汇编代码


GODEBUG=schedtrace=1000 go run ./test.go
SCHED 0ms: gomaxprocs=8 idleprocs=6 threads=4 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]

GODEBUG=gctrace=1 go run main.go
GC 时候输出的内容格式如下：
gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
格式解释说明如下：

gc #：GC 编号，每次 GC 时递增
@#s：程序自启动以来的时间（单位秒）
#%：程序自启动以来花费在 GC 上的时间百分比
#+…+#：GC 各阶段花费的时间，分别为单个P的墙上时间和累计CPU时间
#->#-># MB：分别表示 GC 启动时, GC 结束时, GC 活动时的堆大小
#MB goal：下一次触发 GC 的内存占用阈值
#P：当前使用的处理器P的数量
gc 100 @0.904s 11%: 0.043+2.8+0.029 ms clock, 0.34+3.4/5.4/13.6+0.23 ms cpu, 10->11->6 MB, 12 MB goal, 8 P

gc 100：第 100 次 GC
@0.904s：当前时间是程序启动后的0.904s
11%：程序启动后到现在共花费 11% 的时间在 GC 上
0.043+2.8+0.029 ms clock
0.043：表示单个 P 在 mark 阶段的 STW 时间
2.8：表示所有 P 的 concurrent mark（并发标记）所使用的时间
0.029：表示单个 P 的 markTermination 阶段的 STW 时间
0.34+3.4/5.4/0+0.23 ms cpu
0.34：表示整个进程在 mark 阶段 STW 停顿的时间，一共0.34秒，即 0.043 * 8
3.4/5.4/13.6：3.4 表示 mutator assist 占用的时间，5.4 表示 dedicated + fractional 占用的时间，13.6 表示 idle 占用的时间。这三块累计时间为22.4，即2.8 * 8
0.23 ms：0.23 表示整个进程在 markTermination 阶段 STW 时间，即0.029 * 8
10->11->6 MB
10：表示开始 mark 阶段前的 heap_live 大小
11：表示开始 markTermination 阶段前的 heap_live 大小
6：表示被标记对象的大小
12 MB goal：表示下一次触发 GC 回收的阈值是 12 MB
8 P：本次 GC 一共涉及8 个P
