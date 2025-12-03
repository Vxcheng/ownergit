### PProf

想要进行性能优化，首先瞩目在 Go 自身提供的工具链来作为分析依据，本文将带你学习、使用 Go 后花园，涉及如下：

- runtime/pprof：采集程序（非 Server）的运行数据进行分析
- net/http/pprof：采集 HTTP Server 的运行时数据进行分析

### 是什么

pprof 是用于可视化和分析性能分析数据的工具

pprof 以 [profile.proto](https://github.com/google/pprof/blob/master/proto/profile.proto) 读取分析样本的集合，并生成报告以可视化并帮助分析数据（支持文本和图形报告）

profile.proto 是一个 Protocol Buffer v3 的描述文件，它描述了一组 callstack 和 symbolization 信息， 作用是表示统计分析的一组采样的调用栈，是很常见的 stacktrace 配置文件格式

### 支持什么使用模式

- Report generation：报告生成
- Interactive terminal use：交互式终端使用
- Web interface：Web 界面

### 可以做什么

- CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置
- Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏
- Block Profiling：阻塞分析，记录 goroutine 阻塞等待同步（包括定时器通道）的位置
- Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况

## 一个简单的例子

我们将编写一个简单且有点问题的例子，用于基本的程序初步分析

### 编写 demo 文件

（1）demo.go，文件内容：

```
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/EDDYCJY/go-pprof-example/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
```

（2）data/d.go，文件内容：

```
package data

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
```

运行这个文件，你的 HTTP 服务会多出 /debug/pprof 的 endpoint 可用于观察应用程序的情况

### 分析

#### 一、通过 Web 界面

查看当前总览：访问 `http://127.0.0.1:6060/debug/pprof/`

```
/debug/pprof/

profiles:
0	block
5	goroutine
3	heap
0	mutex
9	threadcreate

full goroutine stack dump
```

这个页面中有许多子页面，咱们继续深究下去，看看可以得到什么？

- cpu（CPU Profiling）: `$HOST/debug/pprof/profile`，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
- block（Block Profiling）：`$HOST/debug/pprof/block`，查看导致阻塞同步的堆栈跟踪
- goroutine：`$HOST/debug/pprof/goroutine`，查看当前所有运行的 goroutines 堆栈跟踪
- heap（Memory Profiling）: `$HOST/debug/pprof/heap`，查看活动对象的内存分配情况
- mutex（Mutex Profiling）：`$HOST/debug/pprof/mutex`，查看导致互斥锁的竞争持有者的堆栈跟踪
- threadcreate：`$HOST/debug/pprof/threadcreate`，查看创建新 OS 线程的堆栈跟踪

#### 二、通过交互式终端使用

（1）go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60

```
$ go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60

Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile?seconds=60
Saved profile in /Users/eddycjy/pprof/pprof.samples.cpu.007.pb.gz
Type: cpu
Duration: 1mins, Total samples = 26.55s (44.15%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

执行该命令后，需等待 60 秒（可调整 seconds 的值），pprof 会进行 CPU Profiling。结束后将默认进入 pprof 的交互式命令模式，可以对分析的结果进行查看或导出。具体可执行 `pprof help` 查看命令说明

```
(pprof) top10
Showing nodes accounting for 25.92s, 97.63% of 26.55s total
Dropped 85 nodes (cum <= 0.13s)
Showing top 10 nodes out of 21
      flat  flat%   sum%        cum   cum%
    23.28s 87.68% 87.68%     23.29s 87.72%  syscall.Syscall
     0.77s  2.90% 90.58%      0.77s  2.90%  runtime.memmove
     0.58s  2.18% 92.77%      0.58s  2.18%  runtime.freedefer
     0.53s  2.00% 94.76%      1.42s  5.35%  runtime.scanobject
     0.36s  1.36% 96.12%      0.39s  1.47%  runtime.heapBitsForObject
     0.35s  1.32% 97.44%      0.45s  1.69%  runtime.greyobject
     0.02s 0.075% 97.51%     24.96s 94.01%  main.main.func1
     0.01s 0.038% 97.55%     23.91s 90.06%  os.(*File).Write
     0.01s 0.038% 97.59%      0.19s  0.72%  runtime.mallocgc
     0.01s 0.038% 97.63%     23.30s 87.76%  syscall.Write
```

- flat：给定函数上运行耗时
- flat%：同上的 CPU 运行耗时总比例
- sum%：给定函数累积使用 CPU 总比例
- cum：当前函数加上它之上的调用运行总耗时
- cum%：同上的 CPU 运行耗时总比例

最后一列为函数名称，在大多数的情况下，我们可以通过这五列得出一个应用程序的运行情况，加以优化 🤔

（2）go tool pprof http://localhost:6060/debug/pprof/heap

```
$ go tool pprof http://localhost:6060/debug/pprof/heap
Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
Saved profile in /Users/eddycjy/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.008.pb.gz
Type: inuse_space
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 837.48MB, 100% of 837.48MB total
      flat  flat%   sum%        cum   cum%
  837.48MB   100%   100%   837.48MB   100%  main.main.func1
```

- -inuse_space：分析应用程序的常驻内存占用情况
- -alloc_objects：分析应用程序的内存临时分配情况

（3） go tool pprof http://localhost:6060/debug/pprof/block

（4） go tool pprof http://localhost:6060/debug/pprof/mutex

#### 三、PProf 可视化界面

这是令人期待的一小节。在这之前，我们需要简单的编写好测试用例来跑一下

##### 编写测试用例

（1）新建 data/d_test.go，文件内容：

```
package data

import "testing"

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
```

（2）执行测试用例

```
$ go test -bench=. -cpuprofile=cpu.prof
pkg: github.com/EDDYCJY/go-pprof-example/data
BenchmarkAdd-4   	10000000	       187 ns/op
PASS
ok  	github.com/EDDYCJY/go-pprof-example/data	2.300s
```

-memprofile 也可以了解一下

##### 启动 PProf 可视化界面

###### 方法一：

```
$ go tool pprof -http=:8080 cpu.prof
```

###### 方法二：

```
$ go tool pprof cpu.prof
$ (pprof) web
```

如果出现 `Could not execute dot; may need to install graphviz.`，就是提示你要安装 `graphviz` 了 （请右拐谷歌）

##### 查看 PProf 可视化界面

（1）Top

[![image](https://camo.githubusercontent.com/7769ebd53e720bb8127fcd706094a27aa527d5e484fbd7dfae7b201d196a7562/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c7359442e6a7067)](https://camo.githubusercontent.com/7769ebd53e720bb8127fcd706094a27aa527d5e484fbd7dfae7b201d196a7562/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c7359442e6a7067)

（2）Graph

[![image](https://camo.githubusercontent.com/8c4bb159431ddfc5faf43466903f404ea2b9f45d71f2aa275094760092550c5a/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c676c642e6a7067)](https://camo.githubusercontent.com/8c4bb159431ddfc5faf43466903f404ea2b9f45d71f2aa275094760092550c5a/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c676c642e6a7067)

框越大，线越粗代表它占用的时间越大哦

（3）Peek

[![image](https://camo.githubusercontent.com/b161860d96f5e0dcee046c18f00912ae0dd6a2f0e81b2de53ac55fa81614f494/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c524f492e6a7067)](https://camo.githubusercontent.com/b161860d96f5e0dcee046c18f00912ae0dd6a2f0e81b2de53ac55fa81614f494/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c524f492e6a7067)

（4）Source

[![image](https://camo.githubusercontent.com/3df40740c63a16c93f30be4e1ec281d18f4fa1ee2a3bb0ff655c3d22d0017801/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c3454662e6a7067)](https://camo.githubusercontent.com/3df40740c63a16c93f30be4e1ec281d18f4fa1ee2a3bb0ff655c3d22d0017801/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c3454662e6a7067)

通过 PProf 的可视化界面，我们能够更方便、更直观的看到 Go 应用程序的调用链、使用情况等，并且在 View 菜单栏中，还支持如上多种方式的切换

你想想，在烦恼不知道什么问题的时候，能用这些辅助工具来检测问题，是不是瞬间效率翻倍了呢 👌

#### 四、PProf 火焰图

另一种可视化数据的方法是火焰图，需手动安装原生 PProf 工具：

（1） 安装 PProf

```
$ go get -u github.com/google/pprof
```

（2） 启动 PProf 可视化界面:

```
$ pprof -http=:8080 cpu.prof
```

（3） 查看 PProf 可视化界面

打开 PProf 的可视化界面时，你会明显发现比官方工具链的 PProf 精致一些，并且多了 Flame Graph（火焰图）

它就是本次的目标之一，它的最大优点是动态的。调用顺序由上到下（A -> B -> C -> D），每一块代表一个函数，越大代表占用 CPU 的时间更长。同时它也支持点击块深入进行分析！

[![image](https://camo.githubusercontent.com/554fa9a998ad1b155e5c2dc2971b090a61e3b2b09e31a0811fa2144987fafed0/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c6a30302e6a7067)](https://camo.githubusercontent.com/554fa9a998ad1b155e5c2dc2971b090a61e3b2b09e31a0811fa2144987fafed0/68747470733a2f2f73322e617831782e636f6d2f323032302f30322f31352f31786c6a30302e6a7067)

## 获取“炸弹”

炸弹程序的代码我已经放到了 [GitHub](https://blog.wolfogre.com/redirect/v3/A_4-v86v-9Btg9a9FuRKCcgSAwM8Cv46xcU7LxImWv3FQQYW3DshxTsGzDw8cyzMPIIcSogxEgMDPAr-OsXFWhYGO25BBhbcOyH9xTwGTQrFOwbMPDwFzDyCHEqIxQ) 上，你只需要在终端里运行 `go get` 便可获取，注意加上 `-d` 参数，避免下载后自动安装：

```bash
go get -d github.com/wolfogre/go-pprof-practice
cd $GOPATH/src/github.com/wolfogre/go-pprof-practice
```

我们可以简单看一下 `main.go` 文件，里面有几个帮助排除性能调问题的关键的的点，我加上了些注释方便你理解，如下：

```go
package main

import (
	// 略
	_ "net/http/pprof" // 会自动注册 handler 到 http server，方便通过 http 接口获取程序运行采样报告
	// 略
)

func main() {
	// 略

	runtime.GOMAXPROCS(1) // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪

	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	// 略
}
```

除此之外的其他代码你一律不用看，那些都是我为了模拟一个“逻辑复杂”的程序而编造的，其中大多数的问题很容易通过肉眼发现，但我们需要做的是通过 pprof 来定位代码的问题，所以为了保证实验的趣味性请不要提前阅读代码，可以实验完成后再看。

接着我们需要编译一下这个程序并运行，你不用担心依赖问题，这个程序没有任何外部依赖。

```bash
go build
./go-pprof-practice
```

运行后注意查看一下资源是否吃紧，机器是否还能扛得住，坚持一分钟，如果确认没问题，咱们再进行下一步。

控制台里应该会不停的打印日志，都是一些“猫狗虎鼠在不停地吃喝拉撒”的屁话，没有意义，不用细看。

```shell
[root@master go-pprof-practice]# ./go-pprof-practice 
2022/12/14 18:10:33 dog.go:24: dog eat
2022/12/14 18:10:33 dog.go:28: dog drink
2022/12/14 18:10:33 dog.go:32: dog shit
2022/12/14 18:10:33 dog.go:36: dog pee
2022/12/14 18:10:33 dog.go:40: dog run
2022/12/14 18:10:33 dog.go:45: dog howl
2022/12/14 18:10:33 wolf.go:27: wolf eat
2022/12/14 18:10:33 wolf.go:31: wolf drink
2022/12/14 18:10:33 wolf.go:40: wolf shit
2022/12/14 18:10:33 wolf.go:44: wolf pee
2022/12/14 18:10:33 wolf.go:48: wolf run
2022/12/14 18:10:33 wolf.go:52: wolf howl
2022/12/14 18:10:34 cat.go:25: cat eat
2022/12/14 18:10:34 cat.go:29: cat drink
2022/12/14 18:10:34 cat.go:33: cat shit
2022/12/14 18:10:34 cat.go:37: cat pee
2022/12/14 18:10:35 cat.go:43: cat climb
2022/12/14 18:10:35 cat.go:47: cat sneak
```

## 使用 pprof

保持程序运行，打开浏览器访问 `http://localhost:6060/debug/pprof/`，可以看到如下页面：

```web-idl
/debug/pprof/

Types of profiles available:
Count	Profile
8	allocs
3	block
0	cmdline
106	goroutine
8	heap
1	mutex
0	profile
7	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
```

页面上展示了可用的程序运行采样数据，分别有：

| 类型         | 描述                       | 备注                                                                                                                                                                                                                                                             |
| :----------- | :------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| allocs       | 内存分配情况的采样信息     | 可以用浏览器打开，但可读性不高                                                                                                                                                                                                                                   |
| blocks       | 阻塞操作情况的采样信息     | 可以用浏览器打开，但可读性不高                                                                                                                                                                                                                                   |
| cmdline      | 显示程序启动命令及参数     | 可以用浏览器打开，这里会显示 `./go-pprof-practice`                                                                                                                                                                                                             |
| goroutine    | 当前所有协程的堆栈信息     | 可以用浏览器打开，但可读性不高                                                                                                                                                                                                                                   |
| heap         | 堆上内存使用情况的采样信息 | 可以用浏览器打开，但可读性不高                                                                                                                                                                                                                                   |
| mutex        | 锁争用情况的采样信息       | 可以用浏览器打开，但可读性不高                                                                                                                                                                                                                                   |
| profile      | CPU 占用情况的采样信息     | 浏览器打开会下载文件                                                                                                                                                                                                                                             |
| threadcreate | 系统线程创建情况的采样信息 | 可以用浏览器打开，但可读性不高                                                                                                                                                                                                                                   |
| trace        | 程序运行跟踪信息           | 浏览器打开会下载文件，本文不涉及，可另行参阅[《深入浅出 Go trace》](https://blog.wolfogre.com/redirect/v3/AwBGKjtUXC4lQ2UdNqHTCoMSAwM8Cv46xcUtPG6RCPoPbv8CcXH9xQrF_wJJOfr_AlNN-lP_AjMyHP8IQUxTTlFBTjhBFgn-UTESAwM8Cv46xcVaFgY7bkEGFtw7If3FPAZNCsU7Bsw8PAXMPIIcSojF) |

因为 cmdline 没有什么实验价值，trace 与本文主题关系不大，threadcreate 涉及的情况偏复杂，所以这三个类型的采样信息这里暂且不提。除此之外，其他所有类型的采样信息本文都会涉及到，且炸弹程序已经为每一种类型的采样信息埋藏了一个对应的性能问题，等待你的发现。

由于直接阅读采样信息缺乏直观性，我们需要借助 `go tool pprof` 命令来排查问题，这个命令是 go 原生自带的，所以不用额外安装。

我们先不用完整地学习如何使用这个命令，毕竟那太枯燥了，我们一边实战一边学习。

以下正式开始。

## 排查 CPU 占用过高

我们首先通过活动监视器（或任务管理器、top 命令，取决于你的操作系统和你的喜好），查看一下炸弹程序的 CPU 占用：

```
top - 18:19:23 up  2:25,  4 users,  load average: 0.22, 0.36, 0.34
Tasks:   1 total,   0 running,   1 sleeping,   0 stopped,   0 zombie
%Cpu(s): 25.0 us,  0.0 sy,  0.0 ni, 75.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :  7990076 total,   153200 free,  7237052 used,   599824 buff/cache
KiB Swap:        0 total,        0 free,        0 used.   411512 avail Mem 

   PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                                   
 99006 root      20   0 2986816   2.1g   1696 S 106.7 27.1   0:09.78 go-pprof-practi  
```

可以看到 CPU 占用相当高，这显然是有问题的，我们使用 `go tool pprof` 来排场一下：

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

等待一会儿后，进入一个交互式终端：

```
[root@master v2ray-linux-32-v4.21.3]# go tool pprof http://localhost:6060/debug/pprof/profile
Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile
Saved profile in /root/pprof/pprof.go-pprof-practice.samples.cpu.002.pb.gz
File: go-pprof-practice
Type: cpu
Time: Dec 14, 2022 at 6:20pm (CST)
Duration: 30.16s, Total samples = 17.52s (58.10%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
```

 输入 top 命令，查看 CPU 占用较高的调用：

```
(pprof) top
Showing nodes accounting for 17.46s, 99.66% of 17.52s total
Dropped 16 nodes (cum <= 0.09s)
      flat  flat%   sum%        cum   cum%
    16.34s 93.26% 93.26%     16.39s 93.55%  github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Eat
     0.57s  3.25% 96.52%      0.57s  3.25%  runtime.memmove
     0.49s  2.80% 99.32%      0.49s  2.80%  runtime.memclrNoHeapPointers
     0.06s  0.34% 99.66%      1.13s  6.45%  github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Steal
         0     0% 99.66%     16.39s 93.55%  github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Live
         0     0% 99.66%      1.13s  6.45%  github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Live
         0     0% 99.66%     17.52s   100%  main.main
         0     0% 99.66%      1.07s  6.11%  runtime.growslice
         0     0% 99.66%     17.52s   100%  runtime.main
(pprof) 
```

很明显，CPU 占用过高是 `github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Eat` 造成的。

输入 `list Eat`，查看问题具体在代码的哪一个位置：

```
(pprof) list Eat
Total: 17.52s
ROUTINE ======================== github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Eat in /root/test/zuoye/go-pprof-practice/animal/felidae/tiger/tiger.go
    16.34s     16.39s (flat, cum) 93.55% of Total
         .          .     17:   t.Climb()
         .          .     18:   t.Sneak()
         .          .     19:}
         .          .     20:
         .          .     21:func (t *Tiger) Eat() {
         .       10ms     22:   log.Println(t.Name(), "eat")
         .          .     23:   loop := 10000000000
    16.34s     16.38s     24:   for i := 0; i < loop; i++ {
         .          .     25:           // do nothing
         .          .     26:   }
         .          .     27:}
         .          .     28:
         .          .     29:func (t *Tiger) Drink() {
(pprof) 
```

可以看到，是第 24 行那个一百亿次空循环占用了大量 CPU 时间，至此，问题定位成功！

接下来有一个扩展操作：图形化显示调用栈信息，这很酷，但是需要你事先在机器上安装 `graphviz`，大多数系统上可以轻松安装它：

```bash
brew install graphviz # for macos
apt install graphviz # for ubuntu
yum install graphviz # for centos
```

或者你也可以访问 [graphviz 官网](https://blog.wolfogre.com/redirect/v3/A421Yoc_xEV4GG_UO8tV1nMSAwM8Cv46xcU7gjwSbQjbbjsviVpukMUYBkEJFgboxTESAwM8Cv46xcVaFgY7bkEGFtw7If3FPAZNCsU7Bsw8PAXMPIIcSojF)寻找适合自己操作系统的安装方法。

安装完成后，我们继续在上文的交互式终端里输入 `web`，注意，虽然这个命令的名字叫“web”，但它的实际行为是产生一个 .svg 文件，并调用你的系统里设置的默认打开 .svg 的程序打开它。如果你的系统里打开 .svg 的默认程序并不是浏览器（比如可能是你的代码编辑器），这时候你需要设置一下默认使用浏览器打开 .svg 文件，相信这难不倒你。

图中，`tiger.(*Tiger).Eat` 函数的框特别大，箭头特别粗，pprof 生怕你不知道这个函数的 CPU 占用很高，这张图还包含了很多有趣且有价值的信息，你可以多看一会儿再继续。

至此，这一小节使用 pprof 定位 CPU 占用的实验就结束了，你需要输入 `exit` 退出 pprof 的交互式终端。

为了方便进行后面的实验，我们修复一下这个问题，不用太麻烦，注释掉相关代码即可：

```go
func (t *Tiger) Eat() {
	log.Println(t.Name(), "eat")
	//loop := 10000000000
	//for i := 0; i < loop; i++ {
	//	// do nothing
	//}
}
```

之后修复问题的的方法都是注释掉相关的代码，不再赘述。你可能觉得这很粗暴，但要知道，这个实验的重点是如何使用 pprof 定位问题，我们不需要花太多时间在改代码上。

## 排查内存占用过高

重新编译炸弹程序，再次运行，可以看到 CPU 占用率已经下来了，但是内存的占用率仍然很高：

```
[root@master v2ray-linux-32-v4.21.3]# top -p 42476

top - 16:51:35 up 57 min,  4 users,  load average: 0.34, 0.64, 0.68
Tasks:   1 total,   0 running,   1 sleeping,   0 stopped,   0 zombie
%Cpu(s):  4.2 us,  1.9 sy,  0.0 ni, 93.5 id,  0.2 wa,  0.0 hi,  0.3 si,  0.0 st
KiB Mem :  7990076 total,   149484 free,  6819228 used,  1021364 buff/cache
KiB Swap:        0 total,        0 free,        0 used.   837968 avail Mem 

   PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                                   
 42476 root      20   0 3122120   1.7g   1056 S   0.0 22.3   0:03.89 go-pprof-practi                                                           
```

我们再次运行使用 pprof 命令，注意这次使用的 URL 的结尾是 heap：

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

再一次使用 `top`、`list` 来定问问题代码：

```
[root@master v2ray-linux-32-v4.21.3]# go tool pprof http://localhost:6060/debug/pprof/heap
Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
Saved profile in /root/pprof/pprof.go-pprof-practice.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
File: go-pprof-practice
Type: inuse_space
Time: Dec 14, 2022 at 4:51pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1GB, 99.85% of 1GB total
Dropped 17 nodes (cum <= 0.01GB)
      flat  flat%   sum%        cum   cum%
       1GB 99.85% 99.85%        1GB 99.85%  github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Steal
         0     0% 99.85%        1GB 99.85%  github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Live
         0     0% 99.85%        1GB 99.85%  main.main
         0     0% 99.85%        1GB 99.90%  runtime.main
(pprof) list
command list requires an argument
(pprof) list Steal
Total: 1GB
ROUTINE ======================== github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Steal in /root/test/zuoye/go-pprof-practice/animal/muridae/mouse/mouse.go
       1GB        1GB (flat, cum) 99.85% of Total
         .          .     45:
         .          .     46:func (m *Mouse) Steal() {
         .          .     47:   log.Println(m.Name(), "steal")
         .          .     48:   max := constant.Mi
         .          .     49:   for len(m.buffer)*constant.Mi < max {
       1GB        1GB     50:           m.buffer = append(m.buffer, [constant.Mi]byte{})
         .          .     51:   }
         .          .     52:}
(pprof) 
```

可以看到这次出问题的地方在 `github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Steal`，函数内容如下：

```go
func (m *Mouse) Steal() {
	log.Println(m.Name(), "steal")
	max := constant.Gi
	for len(m.buffer) * constant.Mi < max {
		m.buffer = append(m.buffer, [constant.Mi]byte{})
	}
}
```

可以看到，这里有个循环会一直向 m.buffer 里追加长度为 1 MiB 的数组，直到总容量到达 1 GiB 为止，且一直不释放这些内存，这就难怪会有这么高的内存占用了。

现在我们同样是注释掉相关代码来解决这个问题。

再次编译运行，查看内存占用：

```
[root@master v2ray-linux-32-v4.21.3]# top -p 54324

top - 17:06:02 up  1:11,  4 users,  load average: 0.38, 0.29, 0.45
Tasks:   1 total,   0 running,   1 sleeping,   0 stopped,   0 zombie
%Cpu(s):  2.6 us,  1.6 sy,  0.0 ni, 95.5 id,  0.0 wa,  0.0 hi,  0.3 si,  0.0 st
KiB Mem :  7990076 total,  1794180 free,  5080908 used,  1114988 buff/cache
KiB Swap:        0 total,        0 free,        0 used.  2576044 avail Mem 

   PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                                   
 54324 root      20   0  821952  20868   2524 S   3.0  0.3   0:00.78 go-pprof-practi 
```

可以看到内存占用已经将到了 35 MB，似乎内存的使用已经恢复正常，一片祥和。

但是，内存相关的性能问题真的已经全部解决了吗？

## 排查频繁内存回收

你应该知道，频繁的 GC 对 golang 程序性能的影响也是非常严重的。虽然现在这个炸弹程序内存使用量并不高，但这会不会是频繁 GC 之后的假象呢？

为了获取程序运行过程中 GC 日志，我们需要先退出炸弹程序，再在重新启动前赋予一个环境变量，同时为了避免其他日志的干扰，使用 grep 筛选出 GC 日志查看：

```bash
GODEBUG=gctrace=1 ./go-pprof-practice | grep gc
```

日志输出如下：

```
[root@master go-pprof-practice]# GODEBUG=gctrace=1 ./go-pprof-practice |grep gc
gc 1 @0.002s 2%: 0.011+0.36+0.001 ms clock, 0.011+0.14/0.084/0+0.001 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 2 @3.010s 0%: 0.043+0.43+0.001 ms clock, 0.043+0.17/0.17/0+0.001 ms cpu, 17->17->1 MB, 18 MB goal, 1 P
gc 3 @6.017s 0%: 0.028+0.32+0.002 ms clock, 0.028+0.12/0.11/0+0.002 ms cpu, 17->17->1 MB, 18 MB goal, 1 P
gc 4 @9.026s 0%: 0.063+0.34+0.001 ms clock, 0.063+0.11/0.11/0+0.001 ms cpu, 17->17->1 MB, 18 MB goal, 1 P
gc 5 @12.034s 0%: 0.026+0.25+0.001 ms clock, 0.026+0.097/0.085/0+0.001 ms cpu, 17->17->1 MB, 18 MB goal, 1 P
gc 6 @15.044s 0%: 0.028+0.30+0.001 ms clock, 0.028+0.11/0.11/0+0.001 ms cpu, 17->17->1 MB, 18 MB goal, 1 P
gc 7 @18.050s 0%: 0.023+0.31+0.001 ms clock, 0.023+0.11/0.090/0+0.001 ms cpu, 17->17->1 MB, 18 MB goal, 1 P
```

可以看到，GC 差不多每 3 秒就发生一次，且每次 GC 都会从 16MB 清理到几乎 0MB，说明程序在不断的申请内存再释放，这是高性能 golang 程序所不允许的。

如果你希望进一步了解 golang 的 GC 日志可以查看[《如何监控 golang 程序的垃圾回收》](https://blog.wolfogre.com/redirect/v3/A9DNc05mRFLA-ZPsjfPhLuZDu-oKbuLF_wQyMDE2xf8CMDfF_wIwMcUtHy8qzDsGiVTMOxzFMRIDAzwK_jrFxVoWBjtuQQYW3Dsh_cU8Bk0KxTsGzDw8Bcw8ghxKiMU),为保证实验节奏，这里不做展开。

所以接下来使用 pprof 排查时，我们在乎的不是什么地方在占用大量内存，而是什么地方在不停地申请内存，这两者是有区别的。

由于内存的申请与释放频度是需要一段时间来统计的，所有我们保证炸弹程序已经运行了几分钟之后，再运行命令：

```bash
go tool pprof http://localhost:6060/debug/pprof/allocs
```

同样使用 top、list、web 大法：

```

[root@master v2ray-linux-32-v4.21.3]# go tool pprof http://localhost:6060/debug/pprof/allocs
Fetching profile over HTTP from http://localhost:6060/debug/pprof/allocs
Saved profile in /root/pprof/pprof.go-pprof-practice.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
File: go-pprof-practice
Type: alloc_space
Time: Dec 14, 2022 at 5:07pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 17408.70kB, 100% of 17408.70kB total
Showing top 10 nodes out of 17
      flat  flat%   sum%        cum   cum%
   16384kB 94.11% 94.11%    16384kB 94.11%  github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Run (inline)
  512.50kB  2.94% 97.06%   512.50kB  2.94%  runtime.allocm
  512.20kB  2.94%   100%   512.20kB  2.94%  runtime.malg
         0     0%   100%    16384kB 94.11%  github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Live
         0     0%   100%    16384kB 94.11%  main.main
         0     0%   100%    16384kB 94.11%  runtime.main
         0     0%   100%   512.50kB  2.94%  runtime.mstart
         0     0%   100%   512.50kB  2.94%  runtime.mstart0
         0     0%   100%   512.50kB  2.94%  runtime.mstart1
         0     0%   100%   512.50kB  2.94%  runtime.newm
(pprof) list Run
Total: 17MB
ROUTINE ======================== github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Run in /root/test/zuoye/go-pprof-practice/animal/canidae/dog/dog.go
      16MB       16MB (flat, cum) 94.11% of Total
         .          .     38:   log.Println(d.Name(), "pee")
         .          .     39:}
         .          .     40:
         .          .     41:func (d *Dog) Run() {
         .          .     42:   log.Println(d.Name(), "run")
      16MB       16MB     43:   _ = make([]byte, 16*constant.Mi)
         .          .     44:}
         .          .     45:
         .          .     46:func (d *Dog) Howl() {
         .          .     47:   log.Println(d.Name(), "howl")
         .          .     48:}
(pprof)
```

可以看到 `github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Run` 会进行无意义的内存申请，而这个函数又会被频繁调用，这才导致程序不停地进行 GC:

```go
func (d *Dog) Run() {
	log.Println(d.Name(), "run")
	_ = make([]byte, 16 * constant.Mi)
}
```

这里有个小插曲，你可尝试一下将 `16 * constant.Mi` 修改成一个较小的值，重新编译运行，会发现并不会引起频繁 GC，原因是在 golang 里，对象是使用堆内存还是栈内存，由编译器进行逃逸分析并决定，如果对象不会逃逸，便可在使用栈内存，但总有意外，就是对象的尺寸过大时，便不得不使用堆内存。所以这里设置申请 16 MiB 的内存就是为了避免编译器直接在栈上分配，如果那样得话就不会涉及到 GC 了。

我们同样注释掉问题代码，重新编译执行，可以看到这一次，程序的 GC 频度要低很多，以至于短时间内都看不到 GC 日志了：

## 排查协程泄露

由于 golang 自带内存回收，所以一般不会发生内存泄露。但凡事都有例外，在 golang 中，协程本身是可能泄露的，或者叫协程失控，进而导致内存泄露。

我们在浏览器里可以看到，此时程序的协程数已经多达 106 条：

```
/debug/pprof/

Types of profiles available:
Count	Profile
8	allocs
3	block
0	cmdline
106	goroutine
8	heap
1	mutex
0	profile
7	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
```

虽然 106 条并不算多，但对于这样一个小程序来说，似乎还是不正常的。为求安心，我们再次是用 pprof 来排查一下：

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

同样是 top、list、web 大法：

```
(pprof) top
Showing nodes accounting for 105, 99.06% of 106 total
Showing top 10 nodes out of 38
      flat  flat%   sum%        cum   cum%
       104 98.11% 98.11%        104 98.11%  runtime.gopark
         1  0.94% 99.06%          1  0.94%  runtime/pprof.runtime_goroutineProfileWithLabels
         0     0% 99.06%          2  1.89%  bufio.(*Reader).ReadLine
         0     0% 99.06%          2  1.89%  bufio.(*Reader).ReadSlice
         0     0% 99.06%          2  1.89%  bufio.(*Reader).fill
         0     0% 99.06%        100 94.34%  github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Drink.func1
         0     0% 99.06%          1  0.94%  internal/poll.(*FD).Accept
         0     0% 99.06%          2  1.89%  internal/poll.(*FD).Read
         0     0% 99.06%          3  2.83%  internal/poll.(*pollDesc).wait
         0     0% 99.06%          3  2.83%  internal/poll.(*pollDesc).waitRead (inline)
(pprof) list Drink
Total: 106
ROUTINE ======================== github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Drink.func1 in /root/test/zuoye/go-pprof-practice/animal/canidae/wolf/wolf.go
         0        100 (flat, cum) 94.34% of Total
         .          .     29:
         .          .     30:func (w *Wolf) Drink() {
         .          .     31:   log.Println(w.Name(), "drink")
         .          .     32:   for i := 0; i < 10; i++ {
         .          .     33:           go func() {
         .        100     34:                   time.Sleep(30 * time.Second)
         .          .     35:           }()
         .          .     36:   }
         .          .     37:}
         .          .     38:
         .          .     39:func (w *Wolf) Shit() {
(pprof)
```

可能这次问题藏得比较隐晦，但仔细观察还是不难发现，问题在于 `github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Drink` 在不停地创建没有实际作用的协程：

```go
func (w *Wolf) Drink() {
	log.Println(w.Name(), "drink")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}
```

可以看到，Drink 函数每次会释放 10 个协程出去，每个协程会睡眠 30 秒再退出，而 Drink 函数又会被反复调用，这才导致大量协程泄露，试想一下，如果释放出的协程会永久阻塞，那么泄露的协程数便会持续增加，内存的占用也会持续增加，那迟早是会被操作系统杀死的。

我们注释掉问题代码，重新编译运行可以看到，协程数已经降到 4 条了：

```
/debug/pprof/

Types of profiles available:
Count	Profile
8	allocs
3	block
0	cmdline
4	goroutine
8	heap
1	mutex
0	profile
7	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
```

## 排查锁的争用

到目前为止，我们已经解决这个炸弹程序的所有资源占用问题，但是事情还没有完，我们需要进一步排查那些会导致程序运行慢的性能问题，这些问题可能并不会导致资源占用，但会让程序效率低下，这同样是高性能程序所忌讳的。

我们首先想到的就是程序中是否有不合理的锁的争用，我们倒一倒，回头看看上一张图，虽然协程数已经降到 4 条，但还显示有一个 mutex 存在争用问题。

相信到这里，你已经触类旁通了，无需多言，开整。

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

同样是 top、list、web 大法：

```
[root@master v2ray-linux-32-v4.21.3]# go tool pprof http://localhost:6060/debug/pprof/mutex
Fetching profile over HTTP from http://localhost:6060/debug/pprof/mutex
Saved profile in /root/pprof/pprof.go-pprof-practice.contentions.delay.001.pb.gz
File: go-pprof-practice
Type: delay
Time: Dec 14, 2022 at 5:48pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 9.01s, 100% of 9.01s total
      flat  flat%   sum%        cum   cum%
     9.01s   100%   100%      9.01s   100%  sync.(*Mutex).Unlock (inline)
         0     0%   100%      9.01s   100%  github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Howl.func1
(pprof)
```

可以看出来这问题出在 `github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Howl`。但要知道，在代码中使用锁是无可非议的，并不是所有的锁都会被标记有问题，我们看看这个有问题的锁那儿触雷了。

```go
func (w *Wolf) Howl() {
	log.Println(w.Name(), "howl")

	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}
```

可以看到，这个锁由主协程 Lock，并启动子协程去 Unlock，主协程会阻塞在第二次 Lock 这儿等待子协程完成任务，但由于子协程足足睡眠了一秒，导致主协程等待这个锁释放足足等了一秒钟。虽然这可能是实际的业务需要，逻辑上说得通，并不一定真的是性能瓶颈，但既然它出现在我写的“炸弹”里，就肯定不是什么“业务需要”啦。

所以，我们注释掉这段问题代码，重新编译执行，继续。

## 排查阻塞操作

好了，我们开始排查最后一个问题。

在程序中，除了锁的争用会导致阻塞之外，很多逻辑都会导致阻塞。

```
/debug/pprof/

Types of profiles available:
Count	Profile
8	allocs
3	block
0	cmdline
106	goroutine
8	heap
1	mutex
0	profile
7	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
```

可以看到，这里仍有 2 个阻塞操作，虽然不一定是有问题的，但我们保证程序性能，我们还是要老老实实排查确认一下才对。

```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

top、list、web，你懂得。

```
[root@master v2ray-linux-32-v4.21.3]# go tool pprof http://localhost:6060/debug/pprof/block
Fetching profile over HTTP from http://localhost:6060/debug/pprof/block
Saved profile in /root/pprof/pprof.go-pprof-practice.contentions.delay.002.pb.gz
File: go-pprof-practice
Type: delay
Time: Dec 14, 2022 at 5:51pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 140.10s, 100% of 140.10s total
Dropped 4 nodes (cum <= 0.70s)
      flat  flat%   sum%        cum   cum%
    70.05s 50.00% 50.00%     70.05s 50.00%  sync.(*Mutex).Lock (inline)
    70.05s 50.00%   100%     70.05s 50.00%  runtime.chanrecv1
         0     0%   100%     70.05s 50.00%  github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Howl
         0     0%   100%     70.05s 50.00%  github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Live
         0     0%   100%     70.05s 50.00%  github.com/wolfogre/go-pprof-practice/animal/felidae/cat.(*Cat).Live
         0     0%   100%     70.05s 50.00%  github.com/wolfogre/go-pprof-practice/animal/felidae/cat.(*Cat).Pee
         0     0%   100%    140.10s   100%  main.main
         0     0%   100%    140.10s   100%  runtime.main
(pprof)
```

可以看到，阻塞操作位于 `github.com/wolfogre/go-pprof-practice/animal/felidae/cat.(*Cat).Pee`：

```go
func (c *Cat) Pee() {
	log.Println(c.Name(), "pee")

	<-time.After(time.Second)
}
```

你应该可以看懂，不同于睡眠一秒，这里是从一个 channel 里读数据时，发生了阻塞，直到这个 channel 在一秒后才有数据读出，这就导致程序阻塞了一秒而非睡眠了一秒。

这里有个疑点，就是上文中是可以看到有两个阻塞操作的，但这里只排查出了一个，我没有找到其准确原因，但怀疑另一个阻塞操作是程序监听端口提供 porof 查询时，涉及到 IO 操作发生了阻塞，即阻塞在对 HTTP 端口的监听上，但我没有进一步考证。

## 思考题

（1）说说你遇到过哪些性能影响的坑？怎么处理的？
