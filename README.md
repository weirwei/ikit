# ikit
基础功能库

## 功能
- ihttp 请求
- ilog 日志
- iutil 工具包

## 使用
### 安装
```shell
go get -u github.com/weirwei/ikit
```

### ihttp
发送http 请求
目前仅支持post 和get 请求
使用方式见 [测试文件](ihttp/http_test.go)

### ilog
简单封装了`log`
使用时仅直接调用方法即可
```go
ilog.Infof("get a info %s", m)
```

### iutil
网罗了一些好用的工具
