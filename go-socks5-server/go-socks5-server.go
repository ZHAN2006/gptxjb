package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/armon/go-socks5"
)

// 自定义日志写入器，用于记录连接IP地址和数据量
type logWriter struct{}

func (w *logWriter) Write(p []byte) (n int, err error) {
	fields := string(p)
	log.Printf("[信息] %s", fields)
	return len(p), nil
}

func main() {
	// 定义命令行参数
	listenIP := flag.String("s", "0.0.0.0", "监听的IP地址")
	listenPort := flag.Int("l", 8080, "监听的端口")
	showHelp := flag.Bool("h", false, "显示命令行选项帮助信息")

	flag.Parse()

	// 显示帮助信息
	if *showHelp {
		flag.PrintDefaults()
		return
	}

	// 创建 SOCKS5 代理配置
	conf := &socks5.Config{
		Logger: log.New(&logWriter{}, "", 0),
	}

	// 创建 SOCKS5 代理服务器
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("创建 SOCKS5 服务器出错: %v", err)
	}

	// 启动 SOCKS5 代理服务器
	addr := *listenIP + ":" + strconv.Itoa(*listenPort)
	log.Printf("在 %s 上启动 SOCKS5 代理服务器", addr)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		log.Fatalf("启动 SOCKS5 服务器出错: %v", err)
	}
}
