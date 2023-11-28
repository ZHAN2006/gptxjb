<h1>介绍</h1>
socks5转http工具


<h1>编译方法</h1>

Windows 下编译 Windows 64 位版本
```
SET GOOS=windows
SET GOARCH=amd64
go build -o ProxySockhttp.exe
```

Windows 下编译 Linux 64 位版本
```
SET GOOS=linux
SET GOARCH=amd64
go build -o ProxySockhttp
```

Windows 下编译 macOS 64 位版本
```
SET GOOS=darwin
SET GOARCH=amd64
go build -o ProxySockhttp.app
```

<h1>使用方法</h1>

```
ProxySockhttp用法:
  -h    显示帮助信息
  -l string
        连接到的SOCKS5代理服务器的地址 (default ":1080")
  -s string
        代理服务器监听的地址 (default ":8800")
```
