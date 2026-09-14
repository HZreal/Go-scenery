# Go 性能分析简要技术博客：从 pprof 到 trace

## 背景

Go 程序的性能问题通常不会只表现为“慢”。在真实服务里，它可能是 CPU 飙高、内存持续增长、接口延迟抖动、goroutine 堆积、锁竞争、GC 压力变大，或者调度器长时间没有可用执行资源。

`performance/` 目录下的示例围绕 Go 常用性能分析工具展开：

1. `pprof/notServerCase.go`：非 HTTP Server 程序中主动采集 CPU 和内存 profile。
2. `pprof/serverCase.go`：HTTP Server 程序中通过 `net/http/pprof` 暴露分析端点。
3. `goToolTrace/goToolTrace.go`：通过 `runtime/trace` 生成执行 trace。
4. `debug/trace.go`：通过 `GODEBUG=schedtrace` 观察 Go 调度器状态。

本文基于这些文件整理一份后续可复用的 Go 性能分析技术沉淀。

## 一、性能分析先看什么

做 Go 性能分析时，不建议一上来就改代码。更稳妥的方式是先明确问题类型，再选择工具：

| 问题表现 | 首选工具 | 重点观察 |
| --- | --- | --- |
| CPU 使用率高 | `pprof` CPU profile | 热点函数、累计耗时、调用链 |
| 内存增长或 OOM | `pprof` heap profile | 分配来源、存活对象、累计分配 |
| goroutine 堆积 | `pprof` goroutine profile | 阻塞位置、等待原因 |
| 锁竞争严重 | `pprof` mutex/block profile | 锁等待、阻塞耗时 |
| 延迟抖动、调度异常 | `go tool trace` | goroutine 状态、GC、网络、系统调用 |
| 想看调度器概况 | `GODEBUG=schedtrace` | P/M/G、全局队列、本地队列 |

一个通用分析顺序是：

1. 先复现问题，确认是 CPU、内存、阻塞、调度还是外部依赖。
2. 再采集对应 profile 或 trace。
3. 用 `top` 找热点，用调用图或源码视图确认路径。
4. 修改代码后重新采样，对比前后数据。

## 二、pprof：函数级性能画像

`pprof` 适合回答一个核心问题：程序时间或内存主要花在哪里。

Go 中使用 `pprof` 有两种常见方式：

1. 非 Server 程序：代码主动启动和写入 profile。
2. Server 程序：引入 `net/http/pprof`，通过 HTTP 端点采集。

### 1. 非 Server 程序采集 CPU

`pprof/notServerCase.go` 中通过 `runtime/pprof` 主动创建 CPU profile：

```go
cpuFile, err := os.Create(filepath.Join(currentDir, "notServerCpu.prof"))
if err != nil {
    return
}
defer cpuFile.Close()

if err := pprof.StartCPUProfile(cpuFile); err != nil {
    return
}
defer pprof.StopCPUProfile()
```

随后代码循环执行递归版 `fibonacci(35)`，故意制造 CPU 密集型负载：

```go
for i := 0; i < 30; i++ {
    fibonacci(35)
}
```

执行后会生成：

```text
performance/pprof/notServerCpu.prof
```

可以用下面命令查看热点：

```bash
go tool pprof -top performance/pprof/notServerCpu.prof
```

本目录现有 profile 的结果显示，CPU 样本 100% 落在 `main.fibonacci` 上。这说明示例中的性能瓶颈非常明确：递归 Fibonacci 带来了大量重复计算。

```text
flat    flat%   cum     cum%
730ms   100%    730ms   100%   main.fibonacci
```

这里几个字段需要重点理解：

| 字段 | 含义 |
| --- | --- |
| `flat` | 函数自身消耗，不包含它调用的子函数 |
| `flat%` | 函数自身消耗占总采样的比例 |
| `cum` | 函数自身加上子调用的累计消耗 |
| `cum%` | 累计消耗占总采样的比例 |

如果一个函数 `flat` 很高，说明它自身就是热点；如果 `flat` 不高但 `cum` 很高，说明它调用的下游函数很重。

### 2. 非 Server 程序采集内存

同一个文件中还通过 `WriteHeapProfile` 写入内存 profile：

```go
memFile, err := os.Create(filepath.Join(currentDir, "notServerMemory.prof"))
if err != nil {
    return
}
defer memFile.Close()

if err := pprof.WriteHeapProfile(memFile); err != nil {
    return
}
```

查看命令：

```bash
go tool pprof -top performance/pprof/notServerMemory.prof
```

本目录现有内存 profile 显示主要分配来自 `runtime/pprof.StartCPUProfile`。这不是业务内存瓶颈，而是因为示例代码的主要负载是 CPU 递归计算，并没有制造明显的堆对象压力。

这个现象本身也很重要：内存 profile 的结论必须结合业务场景判断。没有真实分配压力时，报告里排第一的函数未必是业务问题。

实际排查内存问题时，可以关注两个维度：

| 维度 | 含义 | 适用场景 |
| --- | --- | --- |
| `inuse_space` | 当前仍存活对象占用的空间 | 查内存泄漏、常驻对象过大 |
| `alloc_space` | 程序累计分配过的空间 | 查频繁分配、GC 压力 |

常用命令：

```bash
go tool pprof -inuse_space -top memory.prof
go tool pprof -alloc_space -top memory.prof
```

### 3. Server 程序使用 net/http/pprof

`pprof/serverCase.go` 中通过匿名导入启用 pprof：

```go
import (
    _ "net/http/pprof"
)
```

当 `http.ListenAndServe` 的 handler 参数为 `nil` 时，pprof 会注册到默认路由上：

```go
log.Fatal(http.ListenAndServe(":28080", nil))
```

访问入口：

```text
http://localhost:28080/debug/pprof/
```

常见端点如下：

| 端点 | 用途 |
| --- | --- |
| `/debug/pprof/profile` | 采集 CPU profile，默认 30 秒 |
| `/debug/pprof/heap` | 查看堆内存 |
| `/debug/pprof/goroutine` | 查看 goroutine 栈 |
| `/debug/pprof/block` | 查看阻塞情况 |
| `/debug/pprof/mutex` | 查看锁竞争 |
| `/debug/pprof/trace` | 采集执行 trace |

采集线上或本地服务的 CPU profile：

```bash
go tool pprof -http=:8888 http://localhost:28080/debug/pprof/profile?seconds=30
```

查看 goroutine：

```bash
go tool pprof -http=:8888 http://localhost:28080/debug/pprof/goroutine
```

### 4. 生产环境使用建议

`serverCase.go` 中的 `RegisterPprof` 展示了单独启动一个 pprof HTTP 服务的方式：

```go
func RegisterPprof() {
    go func() {
        if err := http.ListenAndServe(":28080", nil); err != nil {
            panic("pprof server start error: " + err.Error())
        }
    }()
}
```

生产环境可以采用这种思路，但需要注意：

1. pprof 端口不要直接暴露到公网。
2. 最好绑定内网地址，或放在运维网络、跳板机、服务网格访问控制之后。
3. CPU profile 会带来额外采样成本，采集时间不要过长。
4. heap、goroutine、trace 数据可能包含函数名、路径、请求信息，注意安全边界。
5. 排查完成后保留 profile 文件，便于和优化后的结果对比。

## 三、go tool trace：事件级执行时间线

`pprof` 更偏函数聚合视角，而 `go tool trace` 更偏时间线视角。它能看到 goroutine 在什么时候创建、运行、阻塞、恢复，也能看到 GC、系统调用、网络阻塞等事件。

`goToolTrace/goToolTrace.go` 中通过 `runtime/trace` 生成 trace 文件：

```go
f, err := os.Create("./trace.out")
if err != nil {
    panic(err)
}
defer f.Close()

err = trace.Start(f)
if err != nil {
    panic(err)
}
defer trace.Stop()
```

运行程序后会生成：

```text
performance/goToolTrace/trace.out
```

打开 trace UI：

```bash
go tool trace performance/goToolTrace/trace.out
```

浏览器中常看的几个视图：

| 视图 | 关注点 |
| --- | --- |
| View trace | goroutine 时间线、阻塞、运行状态切换 |
| Goroutine analysis | goroutine 创建、运行、阻塞分布 |
| Network blocking profile | 网络阻塞 |
| Synchronization blocking profile | 同步阻塞 |
| Syscall blocking profile | 系统调用阻塞 |
| Scheduler latency profile | 调度延迟 |

本目录的 trace 示例只打印了一次 `Hello World`，负载很轻，因此调度延迟 profile 中主要能看到 `runtime/trace.Start` 的极少量开销。这说明 trace 更适合放到真实并发、网络、锁竞争或 GC 压力场景中使用。

当问题是“某个接口偶发慢”“goroutine 为什么迟迟没有执行”“GC 期间是否影响延迟”时，trace 往往比单纯 pprof 更直观。

## 四、GODEBUG=schedtrace：观察调度器概况

`debug/trace.go` 是一个简单循环：

```go
for i := 0; i < 5; i++ {
    time.Sleep(time.Second)
    fmt.Println("Hello World")
}
```

编译运行：

```bash
cd performance/debug
go build trace.go
GODEBUG=schedtrace=1000 ./trace
```

示例输出类似：

```text
SCHED 1006ms: gomaxprocs=10 idleprocs=10 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
```

关键字段解释：

| 字段 | 含义 |
| --- | --- |
| `gomaxprocs` | P 的数量，即同时执行 Go 代码的逻辑处理器数量 |
| `idleprocs` | 空闲 P 数量 |
| `threads` | 当前 OS 线程数量 |
| `spinningthreads` | 正在自旋寻找可运行 goroutine 的线程数量 |
| `idlethreads` | 空闲线程数量 |
| `runqueue` | 全局可运行 goroutine 队列长度 |
| `[...]` | 每个 P 的本地运行队列长度 |

如何解读：

1. `runqueue` 长期较高：可能有大量 goroutine 等待调度，CPU 可能成为瓶颈。
2. 本地队列数组中某些值长期较高：可能存在调度不均衡或局部 goroutine 堆积。
3. `idleprocs` 很低且 CPU 很高：说明 P 基本都在执行任务。
4. `threads` 持续增长：需要结合 syscall、cgo、阻塞调用进一步分析。

`schedtrace` 的优势是轻量、直接；缺点是只能看概况，无法替代 pprof 和 trace 的细粒度调用链。

## 五、一次完整的 Go 性能排查流程

可以把性能分析拆成下面几步：

### 1. 明确问题

先确认现象：

- 是 CPU 高，还是内存高？
- 是接口整体慢，还是偶发毛刺？
- 是启动慢、运行一段时间后慢，还是流量上来后慢？
- 是否伴随 goroutine 增长、GC 次数增加、连接池耗尽？

### 2. 采集证据

CPU 问题：

```bash
go tool pprof -http=:8888 http://localhost:28080/debug/pprof/profile?seconds=30
```

内存问题：

```bash
go tool pprof -http=:8888 http://localhost:28080/debug/pprof/heap
```

goroutine 堆积：

```bash
curl -o goroutine.out "http://localhost:28080/debug/pprof/goroutine?debug=2"
```

延迟抖动：

```bash
curl -o trace.out "http://localhost:28080/debug/pprof/trace?seconds=5"
go tool trace trace.out
```

调度器概况：

```bash
GODEBUG=schedtrace=1000 ./your_app
```

### 3. 分析热点

pprof 中先看：

```bash
top
list <func>
web
```

重点判断：

1. 热点是否在业务函数。
2. `flat` 高还是 `cum` 高。
3. 是否存在重复计算、频繁分配、锁等待或阻塞 IO。
4. 优化点是否真的位于请求主路径。

### 4. 修改后对比

优化不能只凭感觉，需要采集优化前后数据：

```bash
go tool pprof -top before.prof
go tool pprof -top after.prof
```

必要时使用 diff：

```bash
go tool pprof -http=:8888 -diff_base before.prof after.prof
```

## 六、常见优化方向

### 1. CPU 热点

如果 CPU 热点集中在计算函数，可以考虑：

1. 减少重复计算，例如缓存中间结果。
2. 使用更合适的算法和数据结构。
3. 降低序列化、反序列化、正则匹配等高成本操作频率。
4. 对可并行任务使用 worker pool，但要控制 goroutine 数量。

`notServerCase.go` 中的递归 Fibonacci 就是典型重复计算案例。真实业务里类似问题可能出现在递归遍历、动态规划缺失、循环中重复解析配置、重复构造对象等场景。

### 2. 内存分配

如果 heap profile 显示某些路径频繁分配，可以考虑：

1. 复用 buffer，例如 `bytes.Buffer`、`sync.Pool`。
2. 减少不必要的字符串拼接和 `[]byte` 转换。
3. 避免在高频循环中创建临时对象。
4. 控制缓存大小，防止无界 map 或 slice 增长。

注意：`sync.Pool` 不是万能优化。它适合临时对象复用，不适合保存有业务生命周期的数据。

### 3. goroutine 堆积

goroutine 数量异常增长时，常见原因包括：

1. channel 无消费者或消费者过慢。
2. 网络请求没有超时。
3. 锁等待或条件变量等待无法释放。
4. 每个请求启动 goroutine，但没有退出条件。
5. 定时任务重复启动，缺少关闭机制。

这类问题优先看 goroutine profile 和 trace。

### 4. 锁竞争

如果怀疑锁竞争，需要在程序中开启 mutex/block profiling：

```go
runtime.SetMutexProfileFraction(1)
runtime.SetBlockProfileRate(1)
```

然后通过 pprof 查看：

```bash
go tool pprof http://localhost:28080/debug/pprof/mutex
go tool pprof http://localhost:28080/debug/pprof/block
```

优化方向包括缩小锁粒度、减少临界区耗时、读多写少场景使用 `sync.RWMutex`、分片锁或无锁数据结构。

## 七、命令速查

非 Server CPU profile：

```bash
go run performance/pprof/notServerCase.go
go tool pprof -top performance/pprof/notServerCpu.prof
go tool pprof -http=:8888 performance/pprof/notServerCpu.prof
```

非 Server memory profile：

```bash
go tool pprof -top performance/pprof/notServerMemory.prof
go tool pprof -http=:8888 performance/pprof/notServerMemory.prof
```

Server pprof：

```bash
go run performance/pprof/serverCase.go
open http://localhost:28080/debug/pprof/
```

采集服务 CPU：

```bash
go tool pprof -http=:8888 "http://localhost:28080/debug/pprof/profile?seconds=30"
```

采集服务 trace：

```bash
curl -o trace.out "http://localhost:28080/debug/pprof/trace?seconds=5"
go tool trace trace.out
```

主动生成 trace：

```bash
cd performance/goToolTrace
go run goToolTrace.go
go tool trace trace.out
```

查看调度器：

```bash
cd performance/debug
go build trace.go
GODEBUG=schedtrace=1000 ./trace
```

## 八、总结

Go 性能分析的关键不是记住某个命令，而是建立一套判断路径：

1. CPU 问题先看 CPU profile。
2. 内存问题先看 heap profile，并区分存活对象和累计分配。
3. goroutine、锁、阻塞问题结合 goroutine、block、mutex profile。
4. 延迟抖动和调度问题使用 `go tool trace`。
5. 调度器总体状态可以用 `GODEBUG=schedtrace` 快速观察。

`performance/` 目录里的示例虽然简单，但覆盖了 Go 性能分析最常用的入口。后续遇到实际性能问题时，可以先按本文的流程采集证据，再结合 profile 和 trace 做定向优化。
