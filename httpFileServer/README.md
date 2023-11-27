<h1>编译方法</h1>

Windows 下编译 Windows 64 位版本
...
SET GOOS=windows
SET GOARCH=amd64
go build -o httpFileServer.exe
...

Windows 下编译 Linux 64 位版本
...
SET GOOS=linux
SET GOARCH=amd64
go build -o httpFileServer
...

Windows 下编译 macOS 64 位版本
...
SET GOOS=darwin
SET GOARCH=amd64
go build -o httpFileServer.app
...

使用方法
...
httpFileServer用法:
  -h                显示帮助
  -p int            指定要侦听的端口（默认值 8081）
  -d string         指定要服务的目录（默认值 .）
...