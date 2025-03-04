# ikit
基础功能库

- [ikit](#ikit)
  - [安装](#安装)
  - [功能](#功能)
  - [使用文档](#使用文档)
    - [igoroutine](#igoroutine)
      - [NewMulti(num int) \*Multi](#newmultinum-int-multi)
      - [(m \*Multi) Run(f func() error)](#m-multi-runf-func-error)
      - [(m \*Multi) Wait() \[\]error](#m-multi-wait-error)
      - [NewDivide(multi \*Multi, opts ...DivideOption) \*Divide](#newdividemulti-multi-opts-divideoption-divide)
      - [OptTotal(total int) DivideOption](#opttotaltotal-int-divideoption)
      - [OptPageSize(pageSize int) DivideOption](#optpagesizepagesize-int-divideoption)
      - [OptPage(page int) DivideOption](#optpagepage-int-divideoption)
      - [(d \*Divide) GetTotal() int](#d-divide-gettotal-int)
      - [(d \*Divide) Run(f func(page, pageSize int) (total int, err error)) \[\]error](#d-divide-runf-funcpage-pagesize-int-total-int-err-error-error)
    - [ihttp](#ihttp)
      - [POST(opt \*Options) (\*Result, error)](#postopt-options-result-error)
      - [GET(opt \*Options) (\*Result, error)](#getopt-options-result-error)
    - [ilog](#ilog)
    - [zlog](#zlog)
      - [InitLog(config Config)](#initlogconfig-config)
      - [CloseLogger()](#closelogger)
      - [Info(ctx \*gin.Context, args ...interface{})](#infoctx-gincontext-args-interface)
      - [Infof(ctx \*gin.Context, format string, args ...interface{})](#infofctx-gincontext-format-string-args-interface)
      - [Debug(ctx \*gin.Context, args ...interface{})](#debugctx-gincontext-args-interface)
      - [Debugf(ctx \*gin.Context, format string, args ...interface{})](#debugfctx-gincontext-format-string-args-interface)
      - [Error(ctx \*gin.Context, args ...interface{})](#errorctx-gincontext-args-interface)
      - [Errorf(ctx \*gin.Context, format string, args ...interface{})](#errorfctx-gincontext-format-string-args-interface)
      - [Warn(ctx \*gin.Context, args ...interface{})](#warnctx-gincontext-args-interface)
      - [Warnf(ctx \*gin.Context, format string, args ...interface{})](#warnfctx-gincontext-format-string-args-interface)
      - [Fatal(ctx \*gin.Context, args ...interface{})](#fatalctx-gincontext-args-interface)
      - [Fatalf(ctx \*gin.Context, format string, args ...interface{})](#fatalfctx-gincontext-format-string-args-interface)
      - [Panic(ctx \*gin.Context, args ...interface{})](#panicctx-gincontext-args-interface)
      - [Panicf(ctx \*gin.Context, format string, args ...interface{})](#panicfctx-gincontext-format-string-args-interface)
    - [iutil](#iutil)
      - [HanLess(s1, s2 string) bool](#hanlesss1-s2-string-bool)
      - [MinInt(a, b int) int](#mininta-b-int-int)
      - [MaxInt(a, b int) int](#maxinta-b-int-int)
      - [LoadYaml(filename, subPath string, s interface{})](#loadyamlfilename-subpath-string-s-interface)
      - [SetRootPath(r string)](#setrootpathr-string)
      - [GetRootPath() string](#getrootpath-string)
      - [ToJson(input interface{}) string](#tojsoninput-interface-string)
      - [Trim(str string) string](#trimstr-string-string)
      - [StringBytes(s string) \[\]byte](#stringbytess-string-byte)
      - [BytesString(b \[\]byte) string](#bytesstringb-byte-string)
      - [StructMap(st interface{}, m map\[string\]interface{}) error](#structmapst-interface-m-mapstringinterface-error)

## 安装
```shell
go get -u github.com/weirwei/ikit
```

## 功能
- ihttp 请求
- ilog 日志
- iutil 工具包

## 使用文档

### igoroutine
对go 的协程进行一些封装，提供便利的工具

- Multi: 协程控制器，可以控制单次任务的最大协程数
- Divide: 分批处理器，可以并发分批处理数据

#### NewMulti(num int) *Multi
> 创建一个Multi，控制协程数为num

---

#### (m *Multi) Run(f func() error)
> Multi 并发执行f()

---

#### (m *Multi) Wait() []error
> 阻塞，等待协程执行完毕，返回错误信息

---

#### NewDivide(multi *Multi, opts ...DivideOption) *Divide
> 创建一个Divide，用于并发分批处理数据

---

#### OptTotal(total int) DivideOption
> 设置数据总量

---

#### OptPageSize(pageSize int) DivideOption
> 设置分页大小

---

#### OptPage(page int) DivideOption
> 设置分页

---

#### (d *Divide) GetTotal() int
> 获取数据总量

---

#### (d *Divide) Run(f func(page, pageSize int) (total int, err error)) []error
> 进行分组运行，入参为执行的函数，返回参数为错误信息

---
---

### ihttp
发送http 请求
目前仅支持post 和get 请求
使用方式见 [测试文件](ihttp/http_test.go)

#### POST(opt *Options) (*Result, error)
> 发送post 请求

---

#### GET(opt *Options) (*Result, error)
> 发送get 请求

---
---

### ilog
简单封装了`log`
使用时仅直接调用方法即可
```go
ilog.Infof("get a info %s", m)
```

---
---

### zlog
基于zap封装的日志库，提供了更便捷的日志记录方式和上下文日志功能

服务间请求的时候，把 `zlog.HeaderKeyLogId` 放入请求头中，即可根据 `logId` 串联日志

#### InitLog(config Config)
> 初始化日志配置

#### CloseLogger()
> 关闭日志，刷新缓冲区

#### Info(ctx *gin.Context, args ...interface{})
> 记录info级别日志

#### Infof(ctx *gin.Context, format string, args ...interface{})
> 格式化记录info级别日志

#### Debug(ctx *gin.Context, args ...interface{})
> 记录debug级别日志

#### Debugf(ctx *gin.Context, format string, args ...interface{})
> 格式化记录debug级别日志

#### Error(ctx *gin.Context, args ...interface{})
> 记录error级别日志

#### Errorf(ctx *gin.Context, format string, args ...interface{})
> 格式化记录error级别日志

#### Warn(ctx *gin.Context, args ...interface{})
> 记录warn级别日志

#### Warnf(ctx *gin.Context, format string, args ...interface{})
> 格式化记录warn级别日志

#### Fatal(ctx *gin.Context, args ...interface{})
> 记录fatal级别日志

#### Fatalf(ctx *gin.Context, format string, args ...interface{})
> 格式化记录fatal级别日志

#### Panic(ctx *gin.Context, args ...interface{})
> 记录panic级别日志

#### Panicf(ctx *gin.Context, format string, args ...interface{})
> 格式化记录panic级别日志

---

### iutil
网罗了一些好用的工具

#### HanLess(s1, s2 string) bool
> 比较两个汉字字符串，如果 `s1 < s2` 返回 `true`
> 
> - 拼音相同，字不同，比较编码大小
> - 字相同，比较下一个字
> - 有字 > 无字
> - 字符串完全相同返回ture

---

#### MinInt(a, b int) int
> 返回较小值

---

#### MaxInt(a, b int) int
> 返回较大值

---

#### LoadYaml(filename, subPath string, s interface{})
> 加载应用根目录（可通过 `SetRootPath(r string)` 设置应用根目录路径）下相对路径为 `subPath` 下的 yaml 文件内容到 `s` 结构体。

---

#### SetRootPath(r string)
> 设置应用根目录路径，默认为 `.`

---

#### GetRootPath() string
> 返回当前应用根目录

---

#### ToJson(input interface{}) string
> 将结构体转化为 json 字符串，如果序列化失败，返回空字符串

---

#### Trim(str string) string
> 去除所有空白字符

---

#### StringBytes(s string) []byte
> 字符串转 byte 数组（使用同一片内存）

---

#### BytesString(b []byte) string
> byte 数组转字符串（使用同一片内存）

---
---


#### StructMap(st interface{}, m map[string]interface{}) error
> 将结构体转 `map`，依赖 json 转换，需要 json tag
