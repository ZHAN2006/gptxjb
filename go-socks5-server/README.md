<h1>介绍</h1>
<p style="font-size: 24px;">简易socks5服务器</p>

<h1>编译方法</h1>

<p>Windows 下编译 Windows 64 位版本</p>
<pre>
<code>
SET GOOS=windows
SET GOARCH=amd64
go build
</code>
</pre>

<p>Windows 下编译 Linux 64 位版本</p>
<pre>
<code>
SET GOOS=linux
SET GOARCH=amd64
go build
</code>
</pre>

<p>Windows 下编译 macOS 64 位版本</p>
<pre>
<code>
SET GOOS=darwin
SET GOARCH=amd64
go build
</code>
</pre>

<h1>使用方法</h1>

<pre>
<code>
 -h    显示命令行选项帮助信息
  -l int
        监听的端口 (default 8080)
  -s string
        监听的IP地址 (default "0.0.0.0")
</code>
</pre>

<h3>注</h3>
<p>此项目使用了<a href="https://github.com/armon/go-socks5">go-socks5</a>为基本</p>
