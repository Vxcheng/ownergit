# 书籍

基础语法	《Go 语言圣经》、A Tour of Go
并发编程	《Go 并发编程实战》、GMP 调度解析
性能优化	《深入理解 Go 语言》、GC 三色标记5
微服务架构	 KiteX、Kratos、Go-zero、gRPC 实践9
开源贡献	参与 Etcd、TiKV、Dtm 等项目的 Issue/PR

《Mastering Go》 [https://github.com/hantmac/Mastering_Go_ZH_CN](https://github.com/hantmac/Mastering_Go_ZH_CN)
《Concurrency in Go》 [https://github.com/hapi666/GOBook/tree/master](https://github.com/hapi666/GOBook/tree/master)
《100 Go Mistakes and How to Avoid Them》 [https://github.com/teivah/100-go-mistakes](https://github.com/teivah/100-go-mistakes)
《high-performance-go》 [https://github.com/geektutu/high-performance-go](https://github.com/geektutu/high-performance-go)
《Go语言高级编程·第二版》-2025, 豆瓣: [https://book.douban.com/subject/37436371/](https://book.douban.com/subject/37436371/)
《Go语言定制指南》- 2022, 5.4K Star: [https://github.com/chai2010/go-ast-book](https://github.com/chai2010/go-ast-book)
《面向WebAssembly编程》- 2020, 1.4K Star: [https://github.com/3dgen/cppwasm-book](https://github.com/3dgen/cppwasm-book)
《Go语言高级编程》-2019, 19.5K Star: [https://github.com/chai2010/advanced-go-programming-book](https://github.com/chai2010/advanced-go-programming-book)
《WebAssembly标准入门》- 2018: [https://github.com/chai2010/wasm-book-code](https://github.com/chai2010/wasm-book-code)

《深入go语言之旅》[https://go.cyub.vip/](https://go.cyub.vip/)

《Golang修养之路》 [https://www.yuque.com/aceld/golang/ithv8f](https://www.yuque.com/aceld/golang/ithv8f)

《幼麟实验室》[https://space.bilibili.com/567195437/lists](https://space.bilibili.com/567195437/lists)

《Go 程序员面试笔试宝典》https://golang.design/go-questions/

《GO 夜读》[https://space.bilibili.com/326749661/lists](https://space.bilibili.com/326749661/lists)

《腾讯 秀才的进阶之路》 https://golangstar.cn

# 博客

[https://github.com/gopherchina](https://github.com/gopherchina)

# Go 进阶

1. 内存逃逸
2. 垃圾回收，标记清理算法
3. 内存模型
4. goroutine调度，GMP模型
5. 性能分析pprof、go tool trace
6. dlv、gdb调试，依赖src源码

# 重点

## 一、GMP 模型核心架构

OS对比

| 维度                 | **进程**                                                | **线程**                                                 | **协程**                                           |
| -------------------- | ------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------- |
| **资源拥有者** | 拥有独立的地址空间、文件描述符等。                            | 共享所属进程的资源。                                           | 共享所属线程的栈，在堆上分配协程栈。                     |
| **调度者**     | **操作系统内核** 。                                     | **操作系统内核** 。                                      | **用户程序或运行时** 。                            |
| **切换开销**   | **最高** 。涉及用户态<->内核态切换、页表刷新、TLB刷新。 | **中等** 。涉及用户态<->内核态切换，但无需刷新地址空间。 | **最低** 。纯用户态操作，只需保存/恢复少量寄存器。 |
| **并发单位**   | 可以独立运行多个程序。                                        | 进程内的多个执行流。                                           | 线程内的多个执行流。                                     |
| **通信方式**   | 进程间通信复杂（管道、消息队列、共享内存等）。                | 可以直接读写共享内存（需加锁同步）。                           | 通过通道（如Go的channel）或共享内存（需加锁）通信。      |
| **栈大小**     | 由内核管理，通常较大。                                        | 通常MB级别，固定分配。                                         | 可动态增长，通常KB级别，初始很小。                       |

![1773912494494](image/go/1773912494494.png)

![1775125802732](image/go/1775125802732.png)

### 1. 三大核心组件

| 组件                    | 说明                       | 数量关系          |
| ----------------------- | -------------------------- | ----------------- |
| **G** (Goroutine) | 轻量级协程，初始栈2KB      | 动态创建          |
| **M** (Machine)   | 操作系统线程，真正执行单元 | 默认上限10000     |
| **P** (Processor) | 逻辑处理器，调度上下文     | 默认等于CPU核心数 |

1. 关键数据结构

```go
// runtime/runtime2.go
type g struct {
    stack       stack   // 协程栈
    sched       gobuf   // 调度上下文
    atomicstatus uint32 // 状态：_Grunnable/_Grunning等
}

type p struct {
    runq     [256]guintptr // 本地运行队列
    runnext  guintptr      // 高优先级G
    m        *m            // 绑定的M
}

type m struct {
    g0      *g       // 调度专用G
    curg    *g       // 当前运行的G
    p       puintptr // 关联的P

}
```

### 2. 工作窃取（Work Stealing）

当P的本地队列为空时：

1. 以1/61概率检查全局队列
2. 从其他P的本地队列窃取一半G
3. 确保各P负载均衡

### 3. 系统调用优化（hand off移交）

**go**

```
// 网络轮询器（NetPoller）
func netpoll(block bool) gList {
    // 使用epoll/kqueue/IOCP等系统接口
    // 将IO就绪的G加入可运行队列
}
```

### 4. 抢占式调度实现

1. **监控线程** （sysmon）每10ms检测运行超过10ms的G
2. 向目标M发送 `SIGURG`信号
3. 信号处理函数修改G的上下文，插入调度调用

![1773910552716](image/go/1773910552716.png)

![1773911798516](image/go/1773911798516.png)

![1773912035256](image/go/1773912035256.png)

### 5.核心设计

![1775210352618](image/go/1775210352618.png)

![1775210421763](image/go/1775210421763.png)

![1775210467031](image/go/1775210467031.png)

![1775210680726](image/go/1775210680726.png)

## 二、GC算法

### 1. 历程

- **Go V1.3之前的标记-清除(mark and sweep)算法**--STW，stop the world；让程序暂停，程序出现卡顿 ****(重要问题)****
- **Go V1.5的三色并发标记法（白清理->灰->黑）--**开始三色标记之前就会加上STW，在扫描确定黑白对象之后再放开STW****
- **Go V1.5的三色标记为什么需要STW-- 否则**对象4合法引用的对象3，却被GC给“误杀”回收
- **Go V1.5的三色标记为什么需要屏障机制(“强-弱” 三色不变式、插入屏障、删除屏障 )**
- **Go V1.8混合写屏障机制**
- **Go V1.8混合写屏障机制的全场景分析**

### 2.概念

**强三色不变式：**不存在黑色对象引用到白色对象的指针**。**

**弱三色不变式：**所有被黑色对象引用的白色对象都处于灰色保护状态。****

**插入屏障：** 在A对象引用B对象的时候，B对象被标记为灰色，**强三色不变式，**仅使用在堆空间对象的操作中；最后再次**对栈重新进行三色标记扫描STW**

**删除屏障: **被删除的对象，如果自身为灰色或者白色，那么被标记为灰色,**弱三色不变式******

混合写屏障（堆区，栈上无屏障）,**弱三色不变式**

**1、GC开始将栈上的对象全部扫描并标记为黑色(之后不再进行第二次重复扫描，无需STW)，**

**2、GC期间，任何在栈上创建的新对象，均为黑色。**

**3、被删除的对象标记为灰色。**

**4、被添加的对象标记为灰色。**

总结：

**GoV1.3- 普通标记清除法，整体过程需要启动STW，效率极低。**

**GoV1.5- 三色标记法， 堆空间启动写屏障，栈空间不启动，全部扫描之后，需要重新扫描一次栈(需要STW)，效率普通**

**GoV1.8-三色标记法，混合写屏障机制， 栈空间不启动，堆空间启动。整个过程几乎不需要STW，效率较高。**

![1774000861363](image/go/1774000861363.png)

![1774083884461](image/go/1774083884461.png)

![1774084305384](image/go/1774084305384.png)

## 三、内存管理

### TCMalloc

page、span、size Class

ThreadCache（挂着FreeList）->CentralCache->PageHeap

tiny->page->mSpan->size Class(Object Size)

MCache->MCentral->MHeap，绑定在P

TCMalloc的内存分离

| **对象**   | **容量**         |
| ---------------- | ---------------------- |
| **小对象** | **(0,256KB]**    |
| **中对象** | **(256KB, 1MB]** |
| **大对象** | **(1MB, +**∞)   |

![1774096880650](image/go/1774096880650.png)

![1774097099510](image/go/1774097099510.png)

![1774097450108](image/go/1774097450108.png)

![1774097431604](image/go/1774097431604.png)

### Go内存模型

![1774097686653](image/go/1774097686653.png)

![1774097940431](image/go/1774097940431.png)

![1774098264093](image/go/1774098264093.png)

![1774098570613](image/go/1774098570613.png)

## 四、关键字

### 1. map

总结：演变历程一览表
时期	核心实现	主要优化/解决的问题	遗留/引入的问题
Go 1.0 ~ 1.7	经典链式哈希表	实现简单	1. 内存碎片，缓存不友好
2. 易受哈希碰撞攻击
Go 1.8 ~ 1.17	优化内存布局 + 溢出桶管理	1. Key/Value 分离（省内存，利GC）
2. 溢出桶统一管理（改善局部性）
3. 哈希种子随机化（防攻击）	并发读写 panic 信息不易调试
Go 1.18+	底层算法不变	并发检测提前：更早、更清晰地 panic，易于调试	Map 本身依然非并发安全

哈希取模、链式哈希表、扩容时渐进式迁移

![1774368083906](image/go/1774368083906.png)

![1774366867865](image/go/1774366867865.png)

![1774366844575](image/go/1774366844575.png)

![1774367488883](image/go/1774367488883.png)

![1774367898527](image/go/1774367898527.png)

![1774368451166](image/go/1774368451166.png)

### 2. 切片

![1773222942828](image/go/1773222942828.png)
![1773138698938](image/go/1773138698938.png)
![1773223676286](image/go/1773223676286.png)

### 3. channel

三种状态和三种操作结果

| 操作     | 空值(nil) | 非空已关闭 | 非空未关闭       |
| -------- | --------- | ---------- | ---------------- |
| 关闭     | panic     | panic      | 成功关闭         |
| 发送数据 | 永久阻塞  | panic      | 阻塞或成功发送   |
| 接收数据 | 永久阻塞  | 永不阻塞   | 阻塞或者成功接收 |

![1774599585992](image/go/1774599585992.png)

![1775545330820](image/go/1775545330820.png)

![1774607206732](image/go/1774607206732.png)

![1774369198364](image/go/1774369198364.png)

![1774403564106](image/go/1774403564106.png)

![1775545100709](image/go/1775545100709.png)

![1775544384422](image/go/1775544384422.png)

# 高性能编程

![1774102920217](image/go/1774102920217.png)

# 并发编程

![1774405271315](image/go/1774405271315.png)

![1774405290551](image/go/1774405290551.png)

![1774405306103](image/go/1774405306103.png)

# 语言对比

![1774368818647](image/go/1774368818647.png)

# Gin

![1774407376226](image/skills/1774407376226.png)

![1774407308424](image/skills/1774407308424.png)
