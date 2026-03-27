# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

## 项目概述

**ownergit** 是一个综合性的 Go 语言学习和实验仓库，涵盖从基础语法到微服务架构的完整知识体系。

- **Go 版本**: 1.24.1
- **平台**: Windows
- **仓库类型**: 个人学习知识库，包含多个独立模块

---

## 常用命令

### 测试

```bash
# 运行所有测试（在各个模块目录下）
go test ./... -v

# 运行单个测试文件
go test ./path/to/package -run TestName -v

# 运行带竞态检测的测试
go test ./... -race

# 运行基准测试
go test -bench=. -benchmem

# 运行模糊测试（advance/fuzz 目录）
go test -fuzz=Fuzz -fuzztime=30s
```

### 构建

```bash
# 构建单个模块（在模块目录下）
go build -o bin/app .

# Kratos 项目构建（external_libs/kratos_go/helloworld）
make build          # 编译到 bin/ 目录
make api            # 生成 API proto 文件
make config         # 生成内部 proto 文件
make generate       # 运行 go generate
make all            # 生成所有文件

# Tars 项目构建（external_libs/tars/HelloGo）
make build
```

### 依赖管理

```bash
# 初始化新模块
go mod init ownergit/path/to/module

# 更新依赖
go mod tidy

# 查看依赖
go mod graph
```

### 代码检查

```bash
# 格式化代码
go fmt ./...

# 静态检查
go vet ./...

# 使用 golangci-lint（beego_api 有配置）
golangci-lint run
```

---

## 代码架构

### 整体结构

仓库采用**多模块架构**，每个主要目录都是独立的 Go 模块（30+ 个 go.mod 文件）。主要分为以下几个部分：

#### 1. learning-goSource/ - Go 语言基础

核心语言特性和标准库学习：

- **keyword_go/**: 关键字用法（defer、goto、switch、并发模式）
- **sync_go/**: 并发同步原语（Mutex、RWMutex、WaitGroup、Cond）
- **net_go/**: 网络编程（TCP、UDP、RPC）
- **io_go/**: I/O 操作和文件处理
- **ctx/**: Context 上下文管理
- **time_go/**: 时间和定时器
- **regexp_go/**: 正则表达式
- **unsafe_go/**: 不安全指针操作
- **strings_go/**: 字符串处理
- **crypt/**: 加密算法
- **ast_go/**: 抽象语法树解析
- **gmp/**: GMP 调度模型调试
- **memory_manager/**: 内存逃逸分析
- **embed_go/**: 文件嵌入
- **fmt_go/**: 格式化（独立模块）

#### 2. external_libs/ - 第三方库实践

主流 Go 库的使用示例：

**Web 框架**:
- `gin_web/`: Gin 框架（go.mod: github.com/gin-gonic/gin v1.10.0）
- `beego_api/`: Beego 框架（带 golangci.yaml 配置）
- `kratos_go/helloworld/`: Kratos 微服务框架（带 Makefile）

**RPC 框架**:
- `grpc/`: gRPC + Protobuf
- `tars/HelloGo/`: Tars RPC 框架（带 Makefile）
- `kitex_go/`: Kitex RPC 框架

**数据库与缓存**:
- `gorm/`: GORM ORM 库
- `redis_gomodule/`: Redis 客户端
- `database/`: 数据库操作

**消息队列**:
- `rabbitmq/`: RabbitMQ 客户端/服务端示例

**工具库**:
- `gocron/`: 定时任务（cv1、cv3 版本）
- `ants_pool/`: Goroutine 池
- `lancet/`: 实用工具库
- `colly/`: 网页爬虫
- `gossh/`: SSH 客户端
- `gomail/`: 邮件发送
- `websocket/`: WebSocket 实现
- `gocobra/`: CLI 框架
- `go_prompt/`: 交互式提示
- `gofx/`: 依赖注入

**测试框架**:
- `gomock/`: Mock 框架
- `gomockery/`: Mock 生成器
- `gomonkey/`: Monkey Patching
- `ginkgo/`: BDD 测试框架

#### 3. advance/ - 进阶主题

深入 Go 语言内部机制：

- **escape/**: 内存逃逸分析
- **pprof/**: 性能分析（CPU、内存、goroutine）
- **generics/**: 泛型（独立模块）
- **fuzz/**: 模糊测试（独立模块）
- **go_work/**: Go Workspace 工作区
- **map/**: Map 内部实现和边界情况
- **slice/**: Slice 内部实现
- **channel/**: Channel 模式（fan-out/fan-in）
- **cgo/**: C 语言互操作
- **assemb/**: 汇编语言
- **reflect/**: 反射 API

#### 4. disgin_patterns/ - 设计模式

经典设计模式的 Go 实现：

- **behavioral/**: 行为型模式
  - Strategy（策略）、Template（模板）、State（状态）
  - Decorator（装饰器）、Facade（外观）、Observer（观察者）
- **creation/**: 创建型模式
  - Builder（建造者）
- **structural/**: 结构型模式
  - Bridge（桥接）
- **active/**: Active Object 模式
- **workflow/**: 工作流模式（Worker Pool）

#### 5. arithmetic/ - 算法与数据结构

- **sort/**: 排序算法
- **linked_list/**: 链表和树
- **encryption/**: AES 加密、位运算
- **conversion/**: 进制转换
- **bloom_filter/**: 布隆过滤器
- **bit_test.go**: 位操作

#### 6. face/ - 面试准备

面试和实战问题（独立模块，依赖 etcd、Redis）：

- **leetcode/**: LeetCode 题目（tree、string、plan）
- **micro/limit/**: 微服务限流（Redis 实现）
- **lock/**: 分布式锁（etcd 实现）
- **thread/**: 线程演示
- **search/**: 文件搜索工具

**依赖** (face/go.mod):
```
github.com/gomodule/redigo v1.9.3
go.etcd.io/etcd/client/v3 v3.6.8
golang.org/x/time v0.14.0
```

#### 7. tools/ - 工具

- **excel-generator/**: Excel 文件生成工具（独立模块）

#### 8. books/ - 文档笔记

- **skills.md**: 微服务技能（限流、熔断、降级）
- **go.md**: Go 语言学习资源和进阶主题（GMP 模型、内存逃逸、GC）
- **pprof.md**: 性能分析指南
- **key.md**: 关键概念
- **limux.md**: Linux 笔记
- **image/**: 文档配图

---

## 关键架构说明

### 模块化设计

每个主要功能区域都是独立的 Go 模块，便于：
- 隔离依赖版本
- 独立测试和构建
- 清晰的职责划分

### 测试组织

- **70+ 测试文件**: 分布在各个模块中
- **测试类型**: 单元测试、基准测试、竞态检测、模糊测试
- **测试框架**: 标准 testing 包 + Ginkgo BDD 框架
- **命名规范**: `*_test.go` 文件与源文件同目录

### Kratos 微服务架构 (external_libs/kratos_go/helloworld)

采用 DDD 分层架构：

```
cmd/helloworld/          # 程序入口
internal/
  ├── biz/              # 业务逻辑层
  ├── data/             # 数据访问层
  ├── service/          # 服务层（gRPC/HTTP）
  └── conf/             # 配置定义
api/                    # API 定义（protobuf）
third_party/            # 第三方 proto 文件
```

**构建流程**:
1. `make init` - 安装工具链（protoc-gen-go、wire 等）
2. `make api` - 生成 API 代码（gRPC、HTTP、OpenAPI）
3. `make config` - 生成配置代码
4. `make generate` - 运行 wire 依赖注入
5. `make build` - 编译二进制

### 微服务关键技术 (books/skills.md)

**限流**:
- 固定计数
- 滑动窗口
- 令牌桶
- 漏桶

**熔断**:
- Hystrix 三态机制：关闭 → 打开 → 半开 → 关闭
- 时间窗口内监控失败率和并发量

**降级**:
- 服务降级策略

---

## 开发工作流

### 创建新模块

```bash
# 1. 创建目录
mkdir -p path/to/newmodule

# 2. 初始化模块
cd path/to/newmodule
go mod init ownergit/path/to/newmodule

# 3. 添加代码和测试
# 4. 运行测试
go test -v
```

### 运行特定模块的测试

```bash
# 进入模块目录
cd external_libs/gin_web

# 运行测试
go test ./... -v -race

# 运行基准测试
go test -bench=. -benchmem
```

### 使用 Kratos 开发

```bash
cd external_libs/kratos_go/helloworld

# 修改 proto 文件后重新生成
make api

# 修改业务逻辑后重新构建
make build

# 运行服务
./bin/helloworld -conf ./configs
```

### 性能分析

```bash
# CPU 分析
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# 内存分析
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# 查看逃逸分析
go build -gcflags="-m -m" ./...
```

---

## 重要约定

### 代码组织

- 每个示例尽量保持独立，避免跨模块依赖
- 测试文件与源文件放在同一目录
- 复杂示例使用子目录组织（如 keyword_go/basic、keyword_go/parallel）

### 命名规范

- 包名使用小写，不使用下划线
- 测试函数以 `Test` 开头
- 基准测试以 `Benchmark` 开头
- 示例函数以 `Example` 开头

### 依赖管理

- 使用 `go mod tidy` 保持依赖清洁
- 避免引入不必要的依赖
- 优先使用标准库

---

## 学习路径建议

1. **基础**: learning-goSource/ 目录下的核心概念
2. **进阶**: advance/ 目录的内部机制
3. **实战**: external_libs/ 的第三方库使用
4. **架构**: disgin_patterns/ 的设计模式
5. **面试**: face/ 目录的算法和系统设计

---

## 参考资源

详见 books/go.md 中的书籍和博客列表，包括：
- 《Go 语言圣经》
- 《Go 并发编程实战》
- 《Mastering Go》
- 《100 Go Mistakes and How to Avoid Them》
- GopherChina 社区资源
