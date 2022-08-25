# ikit
基础功能库

* [安装](#安装)
* [功能](#功能)
* [使用文档](#使用文档) 
    * [ihttp](#ihttp)
        * [POST(opt *Options) (*Result, error)](#postopt-options-result-error)
        * [GET(opt *Options) (*Result, error)](#getopt-options-result-error)
    * [ilog](#ilog)
    * [iutil](#iutil)
        * [HanLess(s1, s2 string) bool](#hanlesss1-s2-string-bool)
        * [MinInt(a, b int) int](#mininta-b-int-int)
        * [MaxInt(a, b int) int](#maxinta-b-int-int)
        * [LoadYaml(filename, subPath string, s interface{})](#loadyamlfilename-subpath-string-s-interface)
        * [SetRootPath(r string)](#setrootpathr-string)
        * [GetRootPath() string](#getrootpath-string)
        * [ToJson(input interface{}) string](#tojsoninput-interface-string)
        * [Trim(str string) string](#trimstr-string-string)
        * [StringBytes(s string) []byte](#stringbytess-string-byte)
        * [BytesString(b []byte) string](#bytesstringb-byte-string)
        * [StructMap(st interface{}, m map[string]interface{}) error](#structmapst-interface-m-mapstringinterface-error)

## 安装
```shell
go get -u github.com/weirwei/ikit
```

## 功能
- ihttp 请求
- ilog 日志
- iutil 工具包

## 使用文档

### ihttp
发送http 请求
目前仅支持post 和get 请求
使用方式见 [测试文件](ihttp/http_test.go)

#### POST(opt *Options) (*Result, error)
发送post 请求
---
#### GET(opt *Options) (*Result, error)
发送get 请求
---
---
### ilog
简单封装了`log`
使用时仅直接调用方法即可
```go
ilog.Infof("get a info %s", m)
```

### iutil
网罗了一些好用的工具

#### HanLess(s1, s2 string) bool
比较两个汉字字符串，如果 `s1 < s2` 返回 `true`

- 拼音相同，字不同，比较编码大小
- 字相同，比较下一个字
- 有字 > 无字
- 字符串完全相同返回ture

#### MinInt(a, b int) int
返回较小值

#### MaxInt(a, b int) int
返回较大值

#### LoadYaml(filename, subPath string, s interface{})
加载应用根目录（可通过 `SetRootPath(r string)` 设置应用根目录路径）下相对路径为 `subPath` 下的 yaml 文件内容到 `s` 结构体。

#### SetRootPath(r string)
设置应用根目录路径，默认为 `.`

#### GetRootPath() string
返回当前应用根目录

#### ToJson(input interface{}) string
将结构体转化为 json 字符串，如果序列化失败，返回空字符串

#### Trim(str string) string
去除所有空白字符

#### StringBytes(s string) []byte
字符串转 byte 数组（使用同一片内存）

#### BytesString(b []byte) string
byte 数组转字符串（使用同一片内存）

#### StructMap(st interface{}, m map[string]interface{}) error
将结构体转 `map`，依赖 json 转换，需要 json tag
