<h1>介绍</h1>
简易socks5服务器


<h1>编译方法</h1>

Windows 下编译 Windows 64 位版本
```
SET GOOS=windows
SET GOARCH=amd64
go build
```

Windows 下编译 Linux 64 位版本
```
SET GOOS=linux
SET GOARCH=amd64
go build
```

Windows 下编译 macOS 64 位版本
```
SET GOOS=darwin
SET GOARCH=amd64
go build
```

<h1>使用方法</h1>

```
 -h    显示命令行选项帮助信息
  -l int
        监听的端口 (default 8080)
  -s string
        监听的IP地址 (default "0.0.0.0")

```

<h3>注</h3>
此项目使用了<a href="https://github.com/armon/go-socks5">go-socks5</a>为基本
