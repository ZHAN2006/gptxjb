package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	listenAddr = flag.String("s", ":8800", "代理服务器监听的地址")
	socksAddr  = flag.String("l", ":1080", "连接到的SOCKS5代理服务器的地址")
	showHelp   = flag.Bool("h", false, "显示帮助信息")
)

var (
	socks5proxy proxy.Dialer
	client      *http.Client
)

func newClient(dialer proxy.Dialer) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: func(net, addr string) (net.Conn, error) {
				return socks5proxy.Dial(net, addr)
			},
		},
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -s <代理服务器地址> -l <SOCKS5代理地址> [-h]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	log.Printf("代理服务器正在监听 %s\n", *listenAddr)

	var err error
	socks5proxy, err = proxy.SOCKS5("tcp", *socksAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatal("创建SOCKS5代理失败:", err)
	}
	client = newClient(socks5proxy)

	hndl := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "CONNECT" {
			return
		}

		log.Printf("处理对 %s 的CONNECT请求\n", req.Host)

		serverConn, err := socks5proxy.Dial("tcp", req.Host)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error() + "\n"))
			log.Printf("连接到 %s 失败: %v\n", req.Host, err)
			return
		}
		hijacker, ok := w.(http.Hijacker)
		if !ok {
			serverConn.Close()
			w.WriteHeader(500)
			w.Write([]byte("无法转换为Hijacker\n"))
			log.Println("无法转换为Hijacker")
			return
		}
		w.WriteHeader(200)
		log.Printf("已建立到 %s 的连接\n", req.Host)
		_, bio, err := hijacker.Hijack()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error() + "\n"))
			serverConn.Close()
			log.Printf("无法劫持连接: %v\n", err)
			return
		}
		go io.Copy(serverConn, bio)
		go io.Copy(bio, serverConn)
	})

	err = http.ListenAndServe(*listenAddr, hndl)
	if err != nil {
		log.Fatal("代理服务器启动失败:", err)
	}
}
